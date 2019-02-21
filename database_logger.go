// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

// Success logs a successful event then continues execution
func (d *Database) Success(where string, message string) {
	d.Log.SetPrefix("Success")
	d.Log.Println("|WHERE: ", where, "|MSG: ", message)
}

// Info logs a general information event then continues execution
func (d *Database) Info(where string, message string) {
	d.Log.SetPrefix("Info")
	d.Log.Println("|WHERE: ", where, "|MSG: ", message)
}

// Warn logs a warning event then continues execution
func (d *Database) Warn(where string, message string) {
	d.Log.SetPrefix("Warning")
	d.Log.Println("|WHERE: ", where, "|MSG: ", message)
}

// Error logs a error event then continues execution
func (d *Database) Error(where string, message string, err error) {
	d.Log.SetPrefix("Error")
	d.Log.Println("|WHERE: ", where, "|MSG: ", message, "|ERR: ", err)
}

// Fatal logs a fatal error event then stops execution
func (d *Database) Fatal(where string, message string, err error) {
	d.Log.SetPrefix("Fatal")
	d.Log.Fatal("|WHERE: ", where, "|MSG: ", message, "|ERR: ", err)

}
