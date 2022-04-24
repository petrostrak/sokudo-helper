// sudo lsof -i -P -n | grep LISTEN
// to see running listened connections
package main

import (
	"github.com/petrostrak/sokudo"
	"github.com/petrostrak/sokudo-helper/handlers"
)

type application struct {
	App      *sokudo.Sokudo
	Handlers *handlers.Handlers
}

func main() {
	s := initApplication()
	s.App.ListenAndServe()
}
