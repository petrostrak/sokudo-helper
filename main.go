// sudo lsof -i -P -n | grep LISTEN
// to see running listened connections
package main

import "github.com/petrostrak/sokudo"

type application struct {
	App *sokudo.Sokudo
}

func main() {
	s := initApplication()
	s.App.ListenAndServe()
}
