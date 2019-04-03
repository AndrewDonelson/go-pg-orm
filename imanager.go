// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL
package pgorm

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"os"
	"reflect"
	"strings"
	"time"
)

/********************************************************************
*** ModelDB:IManager Implementation Methods						  ***
********************************************************************/

// Credit: https://gist.github.com/choestelus/8ddcc7106cc247cb5129d4e9c8ba5d64
func (mdb *ModelDB) loadCertificate() error {

	certfilepath := fmt.Sprintf("%s%s", pgORMDestDir, pgORMcrtFile)
	keyfilepath := fmt.Sprintf("%s%s", pgORMDestDir, pgORMkeyFile)

	if !fileExists(certfilepath) || !fileExists(keyfilepath) {
		mdb.generateCertificate(mdb.opts.Addr, pgORMDestDir)
	}
	//cert, err := tls.LoadX509KeyPair("postgresql.crt", "postgresql.key")
	cert, err := tls.LoadX509KeyPair(pgORMcrtFile, pgORMkeyFile)
	if err != nil {
		log.Printf("failed to load client certificate: %v", err)
		return err
	}

	//CAFile := "root.crt"

	CACert, err := ioutil.ReadFile(pgORMcrtFile)
	if err != nil {
		log.Printf("failed to load server certificate: %v", err)
		return err
	}

	CACertPool := x509.NewCertPool()
	CACertPool.AppendCertsFromPEM(CACert)

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            CACertPool,
		InsecureSkipVerify: true,
		// ServerName:         "localhost",
	}

	mdb.opts.TLSConfig = tlsConfig
	// opt := &pg.Options{
	// 	Addr:      "localhost:5432",
	// 	Database:  "postgres",
	// 	User:      "postgres",
	// 	TLSConfig: tlsConfig,
	// }

	return nil
}

// Generate a self-signed X.509 certificate for a TLS server. Outputs to
// 'pgorm-cert.pem' and 'pgorm-key.pem' and will overwrite existing files.
//
//	host       = "", "Comma-separated hostnames and IPs to generate a certificate for"
//	validFrom  = "start-date", "", "Creation date formatted as Jan 1 15:04:05 2011"
//	validFor   = "duration", 365*24*time.Hour, "Duration that certificate is valid for"
//	isCA       = "ca", false, "whether this cert should be its own Certificate Authority"
//	rsaBits    = "rsa-bits", 2048, "Size of RSA key to generate. Ignored if --ecdsa-curve is set"
//	ecdsaCurve = "ecdsa-curve", "", "ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521"
func (mdb *ModelDB) generateCertificate(host, organization string) error {
	//flag.Parse()

	if len(host) == 0 {
		log.Fatalf("Missing required host parameter")
	}

	// if len(pgORMDestDir) == 0 {
	// 	pgORMDestDir = "./"
	// }

	if len(organization) == 0 {
		organization = "ACME Company"
	}

	validFrom := ""
	validFor := 365 * 24 * time.Hour
	isCA := false
	rsaBits := 2048
	ecdsaCurve := ""

	var priv interface{}
	var err error
	switch ecdsaCurve {
	case "":
		priv, err = rsa.GenerateKey(rand.Reader, rsaBits)
	case "P224":
		priv, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized elliptic curve: %q", ecdsaCurve)
		os.Exit(1)
	}
	if err != nil {
		return err
		//log.Fatalf("failed to generate private key: %s", err)
	}

	var notBefore time.Time
	if len(validFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err = time.Parse("Jan 2 15:04:05 2006", validFrom)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse creation date: %s\n", err)
			//os.Exit(1)
			return err
		}
	}

	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		log.Fatalf("failed to generate serial number: %s", err)
		return err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{organization},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	hosts := strings.Split(host, ",")
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	if isCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey(priv), priv)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
		return err
	}

	certfilepath := fmt.Sprintf("%s%s", pgORMDestDir, pgORMcrtFile)
	certOut, err := os.Create(certfilepath)
	if err != nil {
		log.Fatalf("failed to open %s for writing: %s", certfilepath, err)
		return err
	}

	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		log.Fatalf("failed to write data to %s: %s", certfilepath, err)
		return err
	}

	if err := certOut.Close(); err != nil {
		log.Fatalf("error closing %s: %s", certfilepath, err)
		return err
	}
	log.Printf("wrote %s\n", certfilepath)

	keyfilepath := fmt.Sprintf("%s%s", pgORMDestDir, pgORMkeyFile)
	keyOut, err := os.OpenFile(keyfilepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Printf("failed to open %s for writing: %s", keyfilepath, err)
		return err
	}

	if err := pem.Encode(keyOut, pemBlockForKey(priv)); err != nil {
		log.Fatalf("failed to write data to %s: %s", keyfilepath, err)
		return err
	}

	if err := keyOut.Close(); err != nil {
		log.Fatalf("error closing %s: %s", keyfilepath, err)
		return err
	}

	log.Printf("wrote %s\n", keyfilepath)
	return nil
}

// Register adds the values to the models registry
func (mdb *ModelDB) Register(values ...interface{}) error {
	// do not work on the models first, this is like an insurance policy
	// whenever we encounter any error in the values nothing goes into the registry
	models := make(map[string]reflect.Value)
	if len(values) > 0 {
		for _, val := range values {
			rVal := reflect.ValueOf(val)
			if rVal.Kind() == reflect.Ptr {
				rVal = rVal.Elem()
			}
			switch rVal.Kind() {
			case reflect.Struct:
				models[getTypName(rVal.Type())] = reflect.New(rVal.Type())
			default:
				return errors.New("models: models must be structs")
			}
		}
	}

	for k, v := range models {
		mdb.models[k] = v
	}

	return nil
}

// HasTable check has table or not
func (mdb *ModelDB) HasTable(value interface{}) bool {

	qstr := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_schema = 'schema_name' AND table_name = 'table_name');", &value)

	return false
}

// DropTables Drops All Model Database Tables
func (mdb *ModelDB) DropTables() error {
	if mdb.conf.DropTables {
		for _, v := range mdb.models {
			err := mdb.DropTable(v.Interface())
			if err != nil {
				fmt.Println("error", err)
				return err
			}
		}
		fmt.Println("Deleted")
	}
	return nil
}

// AutoMigrateAll runs migrations for all the registered models
func (mdb *ModelDB) AutoMigrateAll() error {
	if !mdb.IsOpen() {
		fmt.Println("Database not open")
	}

	if mdb.conf.Automigrate {
		for _, v := range mdb.models {
			err := mdb.CreateModel(v.Interface())
			if err != nil {
				fmt.Println("Error", err)
				return err
			}
		}
	}
	return nil
}

// IsOpen returns true if the Model has already established connection
// to the database
func (mdb *ModelDB) IsOpen() bool {
	return (mdb.db != nil)
}

// Count returns the number of registered models
func (mdb *ModelDB) Count() int {
	return len(mdb.models)
}
