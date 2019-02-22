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


func NewDatabase(db *pg.DB, log *log.Logger) iDatabase {
	return &Database{
		DB:db,
		Log:log,
	}
}