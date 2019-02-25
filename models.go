package pgorm

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	//"github.com/NlaakStudiosLLC/GoWAF-Framework/logger"
	//"github.com/NlaakStudiosLLC/GoWAF/framework/config"
	//"github.com/jinzhu/gorm"

	// support none, mysql, sqlite3 and postgresql
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/lib/pq"
)

// Model facilitate database interactions, supports postgres, mysql and foundation
type Model struct {
	iDatabase
	models map[string]reflect.Value
	isOpen bool
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

// NewModel returns a new Model without opening database connection
func NewModel() *Model {
	return &Model{
		models: make(map[string]reflect.Value),
	}
}

// Settings: Database
// Database         string `json:"database" yaml:"database" toml:"database" hcl:"database"`
// DatabaseConn     string `json:"database_conn" yaml:"database_conn" toml:"database_conn" hcl:"database_conn"`
// Automigrate      bool   `json:"automigrate" yaml:"automigrate" toml:"automigrate" hcl:"automigrate"`
// DropTables       bool   `json:"droptables" yaml:"droptables" toml:"droptables" hcl:"droptables"`
// NoModel          bool   `json:"no_model" yaml:"no_model" toml:"no_model" hcl:"no_model"`
func (m *Model) OpenWithParams(conn string, migrate bool, drop bool, nomodel bool) error {
	//try and open a connection to the database defined in the config
	db, err := Open()

	if err != nil {
		return err
	}

	//Success we have a database connection
	m.iDatabase = db
	m.isOpen = true
	
	
	return nil
}


// OpenWithConfig opens database connection with the settings found in cfg
func (m *Model) OpenWithConfig(data []byte) error {

	//See if a dabatase was defined in the config
	if len(data) < 5 {
		// Not using a database
		return nil
	}

	db, err := OpenWithOptions(data)

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

//// DropTables Drops All Model Database Tables
func (m *Model) DropTables(drop bool, verbose bool) {
	//if drop {
		for _, v := range m.models {
			err := m.DeleteModel(v.Interface())
			if err != nil  {
				fmt.Println("errror", err)
			}
		}
		fmt.Println("Deleted")
	}
//}

// AutoMigrateAll runs migrations for all the registered models
func (m *Model) AutoMigrateAll(migrate bool) {
	if migrate {
		for _, v := range m.models {
			err := m.CreateModel(v.Interface())
			if err != nil {
				fmt.Println("Error",err)
			}
		}
	}
}

func getTypName(typ reflect.Type) string {
	if typ.Name() != "" {
		return typ.Name()
	}
	split := strings.Split(typ.String(), ".")
	return split[len(split)-1]
}
