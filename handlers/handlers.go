package handlers

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
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

func (h *Handlers) SessioTest(w http.ResponseWriter, r *http.Request) {
	myData := "bar"

	h.App.Session.Put(r.Context(), "foo", myData)

	myValue := h.App.Session.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", myValue)

	if err := h.App.Render.JetPage(w, r, "sessions", vars, nil); err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}

}
