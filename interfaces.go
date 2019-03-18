// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import "github.com/go-pg/pg"

// IDatabase defines the top level Database methods
// Implementation found in individual files
type IDatabase interface {
	IOpen
	IClose
	IManager
	IOptions
	IGet
	IChange
}

// IManager defines the required methods for accessing the database
// Implementation found in imanager.go
// Status: Done
// Test Coverage: Incomplete
type IManager interface {
	LoadCertificate(pgOptions *pg.Options) error
	GenerateCertificate(host, destDir, organization string) error
	Register(values ...interface{}) error
	DropTables() error
	AutoMigrateAll() error
	IsOpen() bool
	Count() int
}

// IOptions ...
// Implementation found in ioptions.go
// Status: Incomplete
// Test Coverage: Incomplete
type IOptions interface {
	EnableSecured() error
	DisableSecured() error
	SetOptions(pgOpts *pg.Options)
	GetOptions() *pg.Options
	DefaultOptions() *pg.Options
}

// IOpen ...
// Implementation found in iopen.go
// Status: Incomplete
// Test Coverage: Incomplete
type IOpen interface {
	Open() error
	OpenWithParams(dbHost, dbUser, dbPass, dbName string) error
	OpenWithConfig(user, database, password string, cfg []byte) error
	OpenWithOptions(pgOpts *pg.Options) error
}

// IClose ...
// Implementation found in iclose.go
// Status: Incomplete
// Test Coverage: Incomplete
type IClose interface {
	Close()
}

// IGet defines the model get related methods
// Implementation found in iget.go
// Status: Incomplete
// Test Coverage: Incomplete
type IGet interface {
	GetModel(model interface{}) error
	GetAllModels(model interface{}) error
	GetWithCondition(model interface{}, condition interface{}, args ...interface{}) error
	GetAllWithCondition(model interface{}, condition interface{}, args ...interface{}) error
}

// IChange defines the model change related methods
// Implementation found in ichange.go
// Status: Incomplete
// Test Coverage: Incomplete
type IChange interface {
	SaveModel(model interface{}) error
	UpdateModel(model interface{}) error
	DeleteModel(model interface{}) error
	CreateModel(model interface{}) error
	DropTable(model interface{}) error
}
