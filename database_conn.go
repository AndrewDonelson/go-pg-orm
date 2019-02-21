// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"github.com/go-pg/pg"
)

// Open
func (d *Database) Open() {
	d.DB = pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer d.Close()
	if d.DB != nil {
		d.Success("Database.Open", "Connected")
	} else {
		d.Fatal("Database.Open", "Failed", nil)
	}
}

// Close
func (d *Database) Close() {
	d.DB.Close()
	d.Info("Database.Close", "Closed")
}
