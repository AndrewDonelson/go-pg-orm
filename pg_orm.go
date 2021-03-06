// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"log"

	"github.com/go-pg/pg"
)

// Database implements the PostgreSQL ORM
type Database struct {
	DB  *pg.DB
	Log *log.Logger
}

// NewDatabase - create instance of Database
func NewDatabase(db *pg.DB, log *log.Logger) IDatabase {
	return &Database{
		DB:  db,
		Log: log,
	}
}


func NewDB(dbHost, dbUser, dbPass, dbName string, secured, migrate, droptables bool, models ...interface{}) (*Model, error) {
	var err error
	mod := NewModel(migrate, droptables)

	err = mod.OpenWithDefault("postgres", "blog", "")
	if err != nil {
		return nil, err
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
