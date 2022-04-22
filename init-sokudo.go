package main

import (
	"log"
	"os"

	"github.com/petrostrak/sokudo"
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
	skd.InfoLog.Println("Debug is set to", skd.Debug)

	app := &application{
		App: skd,
	}

	return app
}
