// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Package pgorm is a Wrapper that simplifies use of Golang ORM with focus on PostgreSQL
// See Documentation at https://godoc.org/github.com/AndrewDonelson/go-pg-orm
package pgorm

import (
	"log"
	"reflect"

	"github.com/go-pg/pg"
)

const (
	pgORMDestDir = "./"             // Default data folder
	pgORMcrtFile = "pgorm-cert.pem" // Certificate filename
	pgORMkeyFile = "pgorm-key.pem"  // Key filename
)

// ModelDB facilitate database interactions, supports postgres, mysql and foundation
type ModelDB struct {
	IDatabase
	db     *pg.DB                   // Database Connection
	conf   *Config                  // go-pg-orm configuration options
	opts   *pg.Options              // Current pg Options
	models map[string]reflect.Value // Hold map of all registered models
	isOpen bool                     // True of the connection to DB is active/open
}

// NewModelDBEnv create a connection to the db with enviroment variables and registers the provided models
func NewModelDBEnv(models ...interface{}) {
	mdb := &ModelDB{}
	err := mdb.OpenWithConfig("", "", "", nil)
	if err != nil {
		log.Fatal("Error with config")
	}

	err = mdb.syncEnv()
	if err != nil {
		log.Fatal("Syncing configuration with Enviroment")
	}
}

// NewModelDBJSON create a connection to the db with a given JSON and registers the provided models
func NewModelDBJSON(models ...interface{}) {

}

// NewModelDBJSON create a connection to the db with a given YAML and registers the provided models
func NewModelDBYAML(models ...interface{}) {

}

// NewModelDBParams prepares a connection to the db with the given parameters. Does NOT open db.
func NewModelDBParams(dbHost, dbUser, dbPass, dbName string, secured, migrate, droptables bool) (*ModelDB, error) {
	var err error
	mdb := NewModel()

	mdb.defaultConfig()
	mdb.conf.Automigrate = migrate
	mdb.conf.DropTables = droptables
	mdb.conf.Secured = secured
	mdb.conf.DatabaseHost = dbHost
	mdb.conf.DatabaseName = dbName
	mdb.conf.DatabaseUser = dbUser
	mdb.conf.DatabasePassword = dbPass

	mdb.defaultOptions()

	//err = mod.Open()
	//if err != nil {
	//	return nil, err
	//}

	if secured {
		err = mdb.loadCertificate()
		if err != nil {
			return nil, err

		}
	}

	return mdb, nil
}

// NewModel returns a new Model without opening database connection
func NewModel() *ModelDB {
	return &ModelDB{
		models: make(map[string]reflect.Value),
	}
}
