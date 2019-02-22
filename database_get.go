// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import "github.com/pkg/errors"

// GetModel attempts retrieve the given model from the database.
func (d *Database) GetModel(model interface{}) error {

	// Select model by given in model properties.
	err := d.DB.Select(model)
	if err != nil {
		return errors.New("Could not get model")
	}
	return nil
}

// GetAllModels attempts retrieve all models based on the given model from the database.
func (d *Database) GetAllModels(model interface{}) error {
	var all []interface{}
	err := d.DB.Model(&all).Select()
	if err != nil {
		return errors.New("Could not get all models")
	}

	return nil
	//d.DB.Set("gorm:auto_preload", true).Order("created_at desc").Find(model)
}

// GetWithCondition attempts retrieve a model based on the given model and condition from the database.
func (d *Database) GetWithCondition(model interface{}, condition interface{}, args ...interface{}) error {
	//TODO check if conditions - string
	if err := d.DB.Model(model).Where(condition.(string), args...).Select(); err != nil {
		return errors.New("Could not get all models")
	}
	return nil
}

// GetAllWithCondition attempts retrieve all models based on the given model and condition from the database.
func (d *Database) GetAllWithCondition(model interface{}, condition interface{}, args ...interface{}) error {
	//TODO check if conditions - string
	var all []interface{}
	if err := d.DB.Model(all).Where(condition.(string), args...).Select(); err != nil {
		return  errors.New("Could not get all models")
	}
	return nil
}

//// GetRowsWithCondition attempts retrieve a model based on the given model and condition from the database.
//func (d *Database) GetRowsWithCondition(model interface{}, condition interface{}, args ...interface{}) error {
//
//	rows := d.DB.Set("gorm:auto_preload", true).Where(condition, args...).Find(model)
//	if rows.RowsAffected != 1 || rows.Error != nil {
//		return errors.New("Could not get model")
//	}
//
//	return nil
//}
