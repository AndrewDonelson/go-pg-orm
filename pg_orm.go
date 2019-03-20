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
	db      *pg.DB                   // DB Connection
	conf    *Config                  // go-pg-orm configuration options
	opts    *pg.Options              // Current Options
	models  map[string]reflect.Value // Hold map of all registered models
	isOpen  bool                     // True of the connection to DB is active/open
	Migrate bool                     // Set to true to auto migrate tables
	Drop    bool                     // Set to true to auto drop tables on each run
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

// NewModelDBParams create a connection to the db with the given parameters and registers the provided models
func NewModelDBParams(dbHost, dbUser, dbPass, dbName string, secured, migrate, droptables bool, models ...interface{}) (*ModelDB, error) {
	var err error
	mod := &ModelDB{
		models:  make(map[string]reflect.Value),
		Migrate: migrate,
		Drop:    droptables,
	}

	err = mod.Open()
	if err != nil {
		return nil, err
	}

	if secured {
		err = mod.loadCertificate()
		if err != nil {
			return nil, err

		}
	}

	//register new model(s)
	err = mod.Register(&models)
	if err != nil {
		return nil, err
	}

	//migrate model
	err = mod.AutoMigrateAll()
	if err != nil {
		return nil, err
	}

	return mod, nil
}

// NewModel returns a new Model without opening database connection
func NewModel(migrate, dropTable bool) *ModelDB {
	return &ModelDB{
		models:  make(map[string]reflect.Value),
		Migrate: migrate,
		Drop:    dropTable,
	}
}
