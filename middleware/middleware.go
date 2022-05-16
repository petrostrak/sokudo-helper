package middleware

import (
	"github.com/petrostrak/sokudo"
	"github.com/petrostrak/sokudo-helper/data"
)

type Middleware struct {
	App    *sokudo.Sokudo
	Models data.Models
}
