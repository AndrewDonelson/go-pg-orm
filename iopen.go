// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"encoding/json"
	"log"
	"os"

	"github.com/go-pg/pg"
)

// Open attempts to open a connection to Postgres using all default postgres & pg.Options
func (mdb *ModelDB) Open() error {
	return nil
}

// OpenWithParams attempts to open a connection to Postgres using default values but allows overriding several common parameters
func (mdb *ModelDB) OpenWithParams(dbHost, dbUser, dbPass, dbName string) error {
	return nil
}

// OpenWithConfig JSON Options must be converted into pg.Options{}, if not - use default options
func (mdb *ModelDB) OpenWithConfig(user, database, password string, cfg []byte) error {
	opts := pg.Options{}
	err := json.Unmarshal(data, &opts)
	if err != nil {
		//connect with default options
		pgDB := pg.Connect(defaultOptions(user, database, password))
		//return NewDatabase(pgDB, log.New(os.Stdout, "", 1)), nil

	}

	pgDB := pg.Connect(&opts)
	return NewDatabase(pgDB, log.New(os.Stdout, "", 1)), nil
}

// OpenWithOptions ...
func (mdb *ModelDB) OpenWithOptions(pgOpts *pg.Options) error {
	return nil
}
