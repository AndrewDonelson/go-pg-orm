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
)

// Open
func (d *Database) Open() {
	d.OpenWithOptions(defaultOptions())

	if d.DB != nil {
		d.Success("Database.Open", "Connected")
	} else {
		d.Fatal("Database.Open", "Failed", nil)
	}
}

// OpenWithOptions
func (d *Database) OpenWithOptions(opts *pg.Options) {
	d.DB = pg.Connect(opts)
	defer d.Close()
}

// Close

func (d *Database) Close() {
	d.DB.Close()
	d.Info("Database.Close", "Closed")
}

func defaultOptions() *pg.Options {
	return &pg.Options{
		User:     "postgres",
		Database: "postgres",

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
