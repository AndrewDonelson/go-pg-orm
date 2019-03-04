// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"crypto/tls"
	"encoding/json"
	"github.com/go-pg/pg"
	"log"
	"os"
	"time"
)



// OpenWithOptions -  Options must be converted into pg.Options{}, if not - use default options
func openWithOptions(user, database string, data []byte) (iDatabase, error) {
	opts := pg.Options{}
	err := json.Unmarshal(data, &opts)
	if err != nil {
		//connect with default options
		pgDB := pg.Connect(defaultOptions(user, database))
		return NewDatabase(pgDB, log.New(os.Stdout, "", 1)), nil
	}

	pgDB := pg.Connect(&opts)
	return NewDatabase(pgDB, log.New(os.Stdout, "", 1)), nil
}

// openWithDefaultOpts -  Options must be converted into pg.Options{}, if not - use default options
func openWithDefaultOpts(user, database string) (iDatabase, error) {
	pgDB := pg.Connect(defaultOptions(user, database))
	return NewDatabase(pgDB, log.New(os.Stdout, "", 1)), nil
}

// Close closes the database client
func (d *Database) Close() {
	err := d.DB.Close()
	if err != nil {
		d.Error("Database.Close", "Can not Close DB", err)
		return
	}

	d.Info("Database.Close", "Closed")
}

func defaultOptions(user, database string) *pg.Options {
	return &pg.Options{
		User:     user,
		Database: database,

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
