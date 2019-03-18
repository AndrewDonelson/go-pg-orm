// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"log"

	"github.com/go-pg/pg/orm"
	"github.com/pkg/errors"
)

// CreateModel attempts add the given model to database.
func (mdb *ModelDB) CreateModel(model interface{}) error {
	err := mdb.db.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		log.Printf("CreateModel", "Model not created", err)
		return errors.New("Could not create model" + err.Error())
	}

	log.Printf("CreateModel", "Model created successfully")
	return nil
}

// SaveModel attempts add the given model to database.
func (mdb *ModelDB) SaveModel(model interface{}) error {
	err := mdb.db.Insert(model)
	if err != nil {
		log.Printf("SaveModel", "Model not saved", err)
		return errors.New("Could not create model")
	}

	log.Printf("SaveModel", "Model saved successfully")
	return nil
}

// UpdateModel attempts update the given model in the database.
func (mdb *ModelDB) UpdateModel(model interface{}) error {
	err := mdb.db.Update(model)
	if err != nil {
		log.Printf("UpdateModel", "Model not updated", err)
		return errors.New("Could not update model")
	}

	log.Printf("UpdateModel", "Model updated successfully")
	return nil
}

// DeleteModel attempts update the given model in the database.
func (mdb *ModelDB) DeleteModel(model interface{}) error {
	err := mdb.db.Delete(model)
	if err != nil {
		log.Printf("DeleteModel", "Model not delete", err)
		return errors.New("Could not delete model")
	}

	log.Printf("DeleteModel", "Model deleted successfully")
	return nil
}

// DropTable drop table from db
func (mdb *ModelDB) DropTable(model interface{}) error {
	err := mdb.db.DropTable(model, &orm.DropTableOptions{true, true})
	if err != nil {
		log.Printf("DropTable", "Model not drop", err)
		return errors.New("Could not drop table")
	}

	log.Printf("DropTable", "Model droped successfully")
	return nil
}
