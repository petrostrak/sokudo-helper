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

func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	if err := h.App.Render.GoPage(w, r, "home", nil); err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}

}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	if err := h.App.Render.JetPage(w, r, "jet-template", nil, nil); err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}

}
