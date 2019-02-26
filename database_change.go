// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"github.com/pkg/errors"
	"github.com/go-pg/pg/orm"
)

// SaveModel attempts add the given model to database.
func (d *Database) CreateModel (model interface{}) error {
	err := d.DB.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		d.Error("CreateModel", "Model not created", err)
		return errors.New("Could not create model"+ err.Error())
	}

	d.Success("CreateModel", "Model created successfully")
	return nil
}


// SaveModel attempts add the given model to database.
func (d *Database) SaveModel(model interface{}) error {
	err := d.DB.Insert(model)
	if err != nil {
		d.Error("SaveModel", "Model not saved", err)
		return errors.New("Could not create model")
	}

	d.Success("SaveModel", "Model saved successfully")
	return d.GetModel(model)
}

// UpdateModel attempts update the given model in the database.
func (d *Database) UpdateModel(model interface{}) error {
	err := d.DB.Update(model)
	if err != nil {
		d.Error("UpdateModel", "Model not updated", err)
		return errors.New("Could not update model")
	}

	d.Success("UpdateModel", "Model updated successfully")
	return nil
}

// DeleteModel attempts update the given model in the database.
func (d *Database) DeleteModel(model interface{}) error {
	err := d.DB.Delete(model)
	if err != nil {
		d.Error("DeleteModel", "Model not delete", err)
		return errors.New("Could not delete model")
	}

	d.Success("DeleteModel", "Model deleted successfully")
	return nil
}

// DropTable drop table
func (d *Database) DropTable(model interface{}) error {
	err := d.DB.DropTable(model, &orm.DropTableOptions{true, true})
	if err != nil {
		d.Error("DropTable", "Model not drop", err)
		return errors.New("Could not drop table")
	}

	d.Success("DropTable", "Model droped successfully")
	return nil
}
