// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/ghodss/yaml"
	"github.com/hashicorp/hcl"

	"github.com/go-pg/pg"
)

var errCfgUnsupported = errors.New("Config file format not supported")

// Config ...
type Config struct {
	DatabaseHost     string `json:"database_host" yaml:"database_host" toml:"database_host" hcl:"database_host"`
	DatabaseName     string `json:"database_name" yaml:"database_name" toml:"database_name" hcl:"database_name"`
	DatabaseUser     string `json:"database_user" yaml:"database_user" toml:"database_user" hcl:"database_user"`
	DatabasePassword string `json:"database_password" yaml:"database_password" toml:"database_password" hcl:"database_password"`
	Automigrate      bool   `json:"automigrate" yaml:"automigrate" toml:"automigrate" hcl:"automigrate"`
	DropTables       bool   `json:"droptables" yaml:"droptables" toml:"droptables" hcl:"droptables"`
	Secured          bool   `json:"secured" yaml:"secured" toml:"secured" hcl:"secured"`
}

// newConfig reads configuration from path. The format is deduced from the file extension
//	* .json    - is decoded as json
//	* .yml     - is decoded as yaml
//	* .toml    - is decoded as toml
//  * .hcl	   - is decoded as hcl
func (mdb *ModelDB) newConfig(path string) (*Config, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	switch filepath.Ext(path) {
	case ".json":
		jerr := json.Unmarshal(data, cfg)
		if jerr != nil {
			return nil, jerr
		}
	case ".toml":
		_, terr := toml.Decode(string(data), cfg)
		if terr != nil {
			return nil, terr
		}
	case ".yml":
		yerr := yaml.Unmarshal(data, cfg)
		if yerr != nil {
			return nil, yerr
		}
	case ".hcl":
		obj, herr := hcl.Parse(string(data))
		if herr != nil {
			return nil, herr
		}
		if herr = hcl.DecodeObject(&cfg, obj); herr != nil {
			return nil, herr
		}
	default:
		return nil, errCfgUnsupported
	}

	// err = mdb.syncEnv(cfg)
	// if err != nil {
	// 	return nil, err
	// }

	return cfg, nil
}

// defaultConfig returns the default configuration settings.
func (mdb *ModelDB) defaultConfig() *Config {
	mdb.conf = &Config{
		DatabaseHost:     "",
		DatabaseName:     "mydb",
		DatabaseUser:     "postgres",
		DatabasePassword: "postgres",
		Automigrate:      true,
		DropTables:       true,
		Secured:          true,
	}
	return mdb.conf
}

// SyncEnv overrides c field's values that are set in the environment.
//
// The environment variable names are derived from config fields by underscoring, and uppercasing
// the name. E.g. AppName will have a corresponding environment variable APP_NAME
//
// NOTE only int, string and bool fields are supported and the corresponding values are set.
// when the field value is not supported it is ignored.
func (mdb *ModelDB) syncEnv() error {
	cfg := reflect.ValueOf(mdb.conf).Elem()
	cTyp := cfg.Type()

	for k := range make([]struct{}, cTyp.NumField()) {
		field := cTyp.Field(k)

		cm := getEnvName(field.Name)
		env := os.Getenv(cm)
		if env == "" {
			continue
		}
		switch field.Type.Kind() {
		case reflect.String:
			cfg.FieldByName(field.Name).SetString(env)
		case reflect.Int:
			v, err := strconv.Atoi(env)
			if err != nil {
				return fmt.Errorf("gowaf: loading config field %s %v", field.Name, err)
			}
			cfg.FieldByName(field.Name).Set(reflect.ValueOf(v))
		case reflect.Bool:
			b, err := strconv.ParseBool(env)
			if err != nil {
				return fmt.Errorf("gowaf: loading config field %s %v", field.Name, err)
			}
			cfg.FieldByName(field.Name).SetBool(b)
		}

	}
	return nil
}

// EnableSecured ...
func (mdb *ModelDB) EnableSecured() error {
	// TODO: Set TLS options for Secured
	return nil
}

// DisableSecured ...
func (mdb *ModelDB) DisableSecured() error {
	// TODO: Set TLS options for Unsecured
	return nil
}

// SetOptions ...
func (mdb *ModelDB) SetOptions(pgOpts *pg.Options) {
	mdb.opts = pgOpts
}

// GetOptions ...
func (mdb *ModelDB) GetOptions() *pg.Options {
	return mdb.opts
}

// defaultOptions ...
func (mdb *ModelDB) defaultOptions() *pg.Options {
	mdb.opts = &pg.Options{
		User:     "postgres",
		Database: "mydb",
		Password: "postgres",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},

		MaxRetries:      1,
		MinRetryBackoff: -1,

		DialTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		PoolSize:           10,
		MaxConnAge:         10 * time.Second,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        10 * time.Second,
		IdleCheckFrequency: 100 * time.Millisecond,
	}

	return mdb.GetOptions()
}
