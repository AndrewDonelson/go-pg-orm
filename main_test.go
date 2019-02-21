package pgorm_test

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/AndrewDonelson/go-pg-orm"
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
		Database: "postgres",

		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},

		MaxRetries:      1,
		MinRetryBackoff: -1,

		DialTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		PoolSize:           10,
		MaxConnAge:         10 * time.Second,
		PoolTimeout:        30 * time.Second,
		IdleTimeout:        10 * time.Second,
		IdleCheckFrequency: 100 * time.Millisecond,
	}
}

func TestRegisterModels(t *testing.T) {

	models := *pgorm.Model.NewModel()
	err = models.Open()
	if err != nil {
		return err
	}

	models.Register(
		&models.User{},
		&models.Story{},
		// ... Register More models here ...
	)
}
func TestDBString(t *testing.T) {
	models = *pgorm.Model.NewModel()
	err = models.Open("")
	if err != nil {
		return err
	}

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

func test_package() {

}
