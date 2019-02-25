// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm


//import "github.com/go-pg/pg"

// IDatabase defines the top level Database methods
type iDatabase interface {
	//IConn
	IGet
	IChange
	ILogger
}

// ILogger defines the logging related methods
type ILogger interface {
	Success(where string, message string)
	Info(where string, message string)
	Warn(where string, message string)
	Error(where string, message string, err error)
	Fatal(where string, message string, err error)
}

// IConn defines the connection related methods
type IConn interface {
	//Open() error
	//OpenWithConnString(conn string) error
	////OpenWithEnv() error
	//OpenWithOptions(opts *pg.Options)
	Close() error
}

// IGet defines the get related methods
type IGet interface {
	GetModel(model interface{}) error
	GetAllModels(model interface{}) error
	GetWithCondition(model interface{}, condition interface{}, args ...interface{}) error
	GetAllWithCondition(model interface{}, condition interface{}, args ...interface{}) error
	//GetRowsWithCondition(interface{}, interface{}, ...interface{}) error
}

// IChange defines the change related methods
type IChange interface {
	SaveModel(model interface{}) error
	UpdateModel(model interface{}) error
	DeleteModel(model interface{}) error
	CreateModel (model interface{}) error
	DropTable(model interface{}) error
}
