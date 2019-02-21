// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import "github.com/pkg/errors"

// SaveModel attempts add the given model to database.
func (d *Database) SaveModel(model interface{}) error {
	err := d.DB.Insert(model)
	if err != nil {
		return errors.New("Could not create model")
	}

	return d.GetModel(model)
}

// UpdateModel attempts update the given model in the database.
func (d *Database) UpdateModel(model interface{}) error {
	err := d.DB.Update(model)
	if err != nil {
		return errors.New("Could not update model")
	}

	return nil
}

// DeleteModel attempts update the given model in the database.
func (d *Database) DeleteModel(model interface{}) error {
	err := d.DB.Delete(model)
	if err != nil {
		return errors.New("Could not delete model")
	}

	return nil
}
