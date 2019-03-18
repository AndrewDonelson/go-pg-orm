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

// DefaultOptions ...
func (mdb *ModelDB) DefaultOptions() *pg.Options {
	mdb.opts = &pg.Options{
		User:     "postgres",
		Database: "mydatabase",
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
