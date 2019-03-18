// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

// // Model facilitate database interactions, supports postgres, mysql and foundation
// type Model struct {
// 	IDatabase
// 	models  map[string]reflect.Value
// 	isOpen  bool
// 	Migrate bool
// 	Drop    bool
// }

// // NewModel returns a new Model without opening database connection
// func NewModel(migrate, dropTable bool) *Model {
// 	return &Model{
// 		models:  make(map[string]reflect.Value),
// 		Migrate: migrate,
// 		Drop:    dropTable,
// 	}
// }

// //OpenWithConfig - opens database connection with the incoming settings,
// //if bad cfg income - use default cfg
// func (m *Model) OpenWithConfig(user, database, password string, cfg []byte) error {
// 	db, err := openWithOptions(user, database, password, cfg)

// 	if err != nil {
// 		return err
// 	}

// 	//Success we have a database connection
// 	m.IDatabase = db
// 	m.isOpen = true
// 	return nil
// }

// //OpenWithConfig - opens database connection with the incoming settings,
// //if bad cfg income - use default cfg
// func (m *Model) OpenWithDefault(user, database, password string) error {
// 	db, err := openWithDefaultOpts(user, database, password)

// 	if err != nil {
// 		return err
// 	}

// 	//Success we have a database connection
// 	m.IDatabase = db
// 	m.isOpen = true
// 	return nil
// }
