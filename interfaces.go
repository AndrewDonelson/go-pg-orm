// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import "github.com/go-pg/pg"

// IDatabase defines the top level Database methods
type IDatabase interface {
	IConn
	IGet
	IChange
	ILogger
}

// ILogger defines the logging related methods
type ILogger interface {
	Success(where string, message string, err error)
	Info(where string, message string, err error)
	Warn(where string, message string, err error)
	Error(where string, message string, err error)
	Fatal(where string, message string, err error)
}

// IConn defines the connection related methods
type IConn interface {
	Open() error
	OpenWithConnString(conn string) error
	OpenWithEnv() error
	OpenWithOptions(opts *pg.Options) error
	Close() error
}

// IGet defines the get related methods
type IGet interface {
	GetModel(interface{}) error
	GetAllModels(interface{})
	GetWithCondition(model interface{}, condition interface{}, args ...interface{}) error
	GetAllWithCondition(model interface{}, condition interface{}, args ...interface{}) error
	GetRowsWithCondition(interface{}, interface{}, ...interface{}) error
}

// IChange defines the change related methods
type IChange interface {
	SaveModel(interface{}) error
	UpdateModel(interface{}) error
	DeleteModel(interface{}) error
}
