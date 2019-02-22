// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"crypto/tls"
	"time"
	"github.com/go-pg/pg"
	"log"
	"os"
)


// Open
func  Open(args ...interface{}) iDatabase{
	if len(args) ==0 {
		db := openWithOptions(defaultOptions())
		return NewDatabase(db, log.New(os.Stdout, "", 1))
	}

	//Todo
	db := openWithOptions(defaultOptions())
	return NewDatabase(db, log.New(os.Stdout, "", 1))
}

// OpenWithOptions
func openWithOptions(opts *pg.Options) *pg.DB{
	pgDB := pg.Connect(opts)
	defer pgDB.Close()
	return  pgDB
}

// Close
func (d *Database) Close() {
	d.DB.Close()
	d.Info("Database.Close", "Closed")
}

func defaultOptions() *pg.Options {
	return &pg.Options{
		User:     "postgres",
		Database: "blog",

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
}