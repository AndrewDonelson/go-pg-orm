// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"reflect"

	"github.com/go-pg/pg"
)

const (
	pgORMDestDir = "./"
	pgORMcrtFile = "pgorm-cert.pem"
	pgORMkeyFile = "pgorm-key.pem"
)

// Model facilitate database interactions, supports postgres, mysql and foundation
type ModelDB struct {
	IDatabase
	db      *pg.DB
	opts    *pg.Options
	models  map[string]reflect.Value
	isOpen  bool
	Migrate bool
	Drop    bool
}

func NewModelDB(dbHost, dbUser, dbPass, dbName string, secured, migrate, droptables bool, models ...interface{}) (*ModelDB, error) {
	var err error
	mod := &Model{
		models:  make(map[string]reflect.Value),
		Migrate: migrate,
		Drop:    dropTable,
	}

	err = mod.OpenWithDefault(dbUser, dbName, dbHost)
	if err != nil {
		return nil, err
	}

	if secured {
		err = mod.LoadCertificate(mod.opts)
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
