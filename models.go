package pgorm

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Model facilitate database interactions, supports postgres, mysql and foundation
type Model struct {
	iDatabase
	models  map[string]reflect.Value
	isOpen  bool
	Migrate bool
	Drop    bool
}

// NewModel returns a new Model without opening database connection
func NewModel(migrate, dropTable bool) *Model {
	return &Model{
		models:  make(map[string]reflect.Value),
		Migrate: migrate,
		Drop:    dropTable,
	}
}

// Register adds the values to the models registry
func (m *Model) Register(values ...interface{}) error {
	// do not work on them.models first, this is like an insurance policy
	// whenever we encounter any error in the values nothing goes into the registry
	models := make(map[string]reflect.Value)
	if len(values) > 0 {
		for _, val := range values {
			rVal := reflect.ValueOf(val)
			if rVal.Kind() == reflect.Ptr {
				rVal = rVal.Elem()
			}
			switch rVal.Kind() {
			case reflect.Struct:
				models[getTypName(rVal.Type())] = reflect.New(rVal.Type())
			default:
				return errors.New("models: models must be structs")
			}
		}
	}

	for k, v := range models {
		m.models[k] = v
	}

	return nil
}

//OpenWithConfig - opens database connection with the incoming settings,
//if bad cfg income - use default cfg
func (m *Model) OpenWithConfig(user, database string, cfg []byte) error {
	db, err := openWithOptions(user, database, cfg)

	if err != nil {
		return err
	}

	//Success we have a database connection
	m.iDatabase = db
	m.isOpen = true
	return nil
}

//OpenWithConfig - opens database connection with the incoming settings,
//if bad cfg income - use default cfg
func (m *Model) OpenWithDefault(user, database string,) error {
	db, err := openWithDefaultOpts(user, database)

	if err != nil {
		return err
	}

	//Success we have a database connection
	m.iDatabase = db
	m.isOpen = true
	return nil
}

// Count returns the number of registered models
func (m *Model) Count() int {
	return len(m.models)
}

// IsOpen returns true if the Model has already established connection
// to the database
func (m *Model) IsOpen() bool {
	return m.isOpen
}

// DropTables Drops All Model Database Tables
func (m *Model) DropTables() error {
	if m.Drop {
		for _, v := range m.models {
			err := m.DropTable(v.Interface())
			if err != nil {
				fmt.Println("error", err)
				return err
			}
		}
		fmt.Println("Deleted")
	}
	return nil
}

// AutoMigrateAll runs migrations for all the registered models
func (m *Model) AutoMigrateAll() error {
	if m.Migrate {
		for _, v := range m.models {
			err := m.CreateModel(v.Interface())
			if err != nil {
				fmt.Println("Error", err)
				return err
			}
		}
	}
	return nil
}

func getTypName(typ reflect.Type) string {
	if typ.Name() != "" {
		return typ.Name()
	}
	split := strings.Split(typ.String(), ".")
	return split[len(split)-1]
}
