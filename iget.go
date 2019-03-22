// Copyright 2019 Andrew Donelson. All rights reserved.
// Use of this source code is governed by a BSD 2-Clause
// "Simplified" License that can be found at
// https://github.com/go-pg/pg/blob/master/LICENSE

// Wrapper that simplifies use of Golang ORM with focus on PostgreSQL

package pgorm

import (
	"log"

	"github.com/pkg/errors"
)

// GetModel attempts retrieve the given model from the database.
func (mdb *ModelDB) GetModel(model interface{}) error {
	// Select model by given in model properties.
	err := mdb.db.Select(model)
	if err != nil {
		log.Println("GetModel", "Model not get", err)
		return errors.New("Could not get model")
	}

	log.Println("GetModel", "Model retrieve successfully")
	return nil
}

// GetAllModels attempts retrieve all models based on the given model from the database.
func (mdb *ModelDB) GetAllModels(model interface{}) error {
	err := mdb.db.Model(model).Select()
	if err != nil {
		log.Println("GetAllModels", "Models not get", err)
		return errors.New("Could not get all models" + err.Error())
	}

	log.Println("GetAllModels", "Model retrieve successfully")
	return nil
}

// GetWithCondition attempts retrieve a model based on the given model and condition from the database.
func (mdb *ModelDB) GetWithCondition(model interface{}, condition interface{}, args ...interface{}) error {
	conditionStr := checkIfString(condition)
	if len(conditionStr) == 0 {
		log.Println("GetWithCondition", "Bad condition")
		return errors.New("Bad condition")
	}

	if err := mdb.db.Model(model).Where(conditionStr, args...).Select(); err != nil {
		log.Println("GetWithCondition", "Models not get", err)
		return errors.New("Could not get all models")
	}

	log.Println("GetWithCondition", "Model retrieve successfully")
	return nil
}

// GetAllWithCondition attempts retrieve all models based on the given model and condition from the database.
func (mdb *ModelDB) GetAllWithCondition(model interface{}, condition interface{}, args ...interface{}) error {
	conditionStr := checkIfString(condition)
	if len(conditionStr) == 0 {
		log.Println("GetWithCondition", "Bad condition")
		return errors.New("Bad condition")
	}

	if err := mdb.db.Model(model).Where(conditionStr, args...).Select(); err != nil {
		log.Println("GetAllWithCondition", "Models not get", err)
		return errors.New("Could not get all models")
	}

	log.Println("GetAllWithCondition", "Model retrieve successfully")
	return nil
}

//Check if conditions - string
func checkIfString(data interface{}) string {
	switch data.(type) {
	case string:
		return data.(string)
	default:
		return ""
	}
}
