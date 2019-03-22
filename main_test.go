package pgorm

import (
	"crypto/tls"
	"log"
	"testing"
	"time"

	"fmt"

	//"github.com/AndrewDonelson/go-pg-orm"
	"github.com/go-pg/pg"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type User struct {
	Id     int64
	Name   string
	Emails []string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *User
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func TestGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "pg")
}

func pgOptions() *pg.Options {
	return &pg.Options{
		User:     "postgres",
		Database: "blog",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxRetries:         1,
		MinRetryBackoff:    -1,
		DialTimeout:        30 * time.Second,
		ReadTimeout:        10 * time.Second,
		WriteTimeout:       10 * time.Second,
		PoolSize:           10,
		MaxConnAge:         10 * time.Second,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        10 * time.Second,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
}

func TestCreateKeyPair(t *testing.T) {
	mdb := ModelDB{}

	err := mdb.generateCertificate("127.0.0.1", "My Company LLC")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterModels(t *testing.T) {

	models := NewModel() //models.Open()

	models.Register(
		&User{},
		&Story{},
		// ... Register More models here ...
	)
}
func TestDBString(t *testing.T) {

	db := pg.Connect(pgOptions())
	defer db.Close()

	wanted := `DB<Addr="localhost:5432">`
	if db.String() != wanted {
		t.Fatalf("got %q, wanted %q", db.String(), wanted)
	}

	db = db.WithParam("param1", "value1").WithParam("param2", 2)
	wanted = `DB<Addr="localhost:5432" param1=value1 param2=2>`
	if db.String() != wanted {
		t.Fatalf("got %q, wanted %q", db.String(), wanted)
	}
}

func TestOnConnect(t *testing.T) {
	opt := pgOptions()
	opt.OnConnect = func(db *pg.Conn) error {
		_, err := db.Exec("SET application_name = 'myapp'")
		return err
	}
	db := pg.Connect(opt)
	defer db.Close()

	var name string
	_, err := db.QueryOne(pg.Scan(&name), "SHOW application_name")
	if err != nil {
		t.Fatal(err)
	}
	if name != "myapp" {
		t.Fatalf(`got %q, wanted "myapp"`, name)
	}
}

func myModels(t *testing.T) []interface{} {
	models := []interface{}{&User{}, &Story{}}
	return models
}

func TestAll(t *testing.T) {
	mdb, err := NewModelDBParams(
		"127.0.0.1", "postgres", "postgres", "mydb",
		true, true, true,
	)

	if err != nil {
		log.Println(err)
		t.Fail()
	}

	err = mdb.Register(&User{}, &Story{})
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	err = mdb.Open()
	if err != nil {
		log.Println(err)
		t.Fail()
	}

	if !mdb.IsOpen() {
		log.Println(err)
		t.Fail()
	}

}
