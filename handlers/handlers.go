package handlers

import (
	"net/http"

	"github.com/petrostrak/sokudo"
)

type Handlers struct {
	App *sokudo.Sokudo
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if err := h.App.Render.Page(w, r, "home", nil, nil); err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}

}
