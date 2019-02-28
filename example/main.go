package main

import (
	goPg "github.com/AndrewDonelson/go-pg-orm"

	"github.com/AndrewDonelson/example/web"
	"github.com/AndrewDonelson/example/model"


)

func main() {
	var err error
	mod := goPg.NewModel(true, true)

	//use default configs - []byte("")
	err = mod.OpenWithConfig([]byte(""))
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
