package main

import (
	"log"
	"os"

	"github.com/petrostrak/sokudo"
	"github.com/petrostrak/sokudo-helper/handlers"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init sokudo
	skd := &sokudo.Sokudo{}
	err = skd.New(path)
	if err != nil {
		log.Fatal(err)
	}

	skd.AppName = "myapp"

	myHandlers := &handlers.Handlers{
		App: skd,
	}

	app := &application{
		App:      skd,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
