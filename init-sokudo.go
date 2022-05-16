package main

import (
	"log"
	"os"

	"github.com/petrostrak/sokudo"
	"github.com/petrostrak/sokudo-helper/data"
	"github.com/petrostrak/sokudo-helper/handlers"
	"github.com/petrostrak/sokudo-helper/middleware"
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

	myMiddleware := &middleware.Middleware{
		App: skd,
	}

	myHandlers := &handlers.Handlers{
		App: skd,
	}

	app := &application{
		App:        skd,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}

	app.App.Routes = app.routes()
	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models
	app.Middleware.Models = app.Models

	return app
}
