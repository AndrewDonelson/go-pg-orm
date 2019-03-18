package main

import (
	goPg "github.com/AndrewDonelson/go-pg-orm"

	"github.com/AndrewDonelson/go-pg-orm/example/model"
	"github.com/AndrewDonelson/go-pg-orm/example/web"
)

const (
	userName = "postgres"
	dbname   = "blog"
	password = ""
)

func main() {
	var err error
	mod, moderr := goPg.NewDB(
		"localhost",      //IP/Domain
		userName,         //DB Username
		password,         //DB Password
		dbname,           //DB Name
		true,             //User TLS Secure
		true,             //Automigrate
		true,             //Drop Tables
		&model.Article{}, //Models to use
	)
	if moderr != nil {
		return
	}
	//mod := goPg.NewModel(true, true)

	err = mod.OpenWithDefault(userName, dbname, password)
	if err != nil {
		return
	}

	//register new model
	err = mod.Register(&model.Article{})
	if err != nil {
		return
	}

	//migrate model
	err = mod.AutoMigrateAll()
	if err != nil {
		return
	}

	web.StartServer(mod)
}
