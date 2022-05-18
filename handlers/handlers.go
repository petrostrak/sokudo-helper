package handlers

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/petrostrak/sokudo"
	"github.com/petrostrak/sokudo-helper/data"
)

type Handlers struct {
	App    *sokudo.Sokudo
	Models data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	if err := h.render(w, r, "home", nil, nil); err != nil {
		h.printError("error rendering:", err)
	}

}

func (h *Handlers) GoPage(w http.ResponseWriter, r *http.Request) {
	if err := h.App.Render.GoPage(w, r, "home", nil); err != nil {
		h.printError("error rendering:", err)
	}

}

func (h *Handlers) JetPage(w http.ResponseWriter, r *http.Request) {
	if err := h.App.Render.JetPage(w, r, "jet-template", nil, nil); err != nil {
		h.printError("error rendering:", err)
	}

}

func (h *Handlers) SessioTest(w http.ResponseWriter, r *http.Request) {
	myData := "bar"

	h.sessionPut(r.Context(), "foo", myData)

	myValue := h.App.Session.GetString(r.Context(), "foo")

	vars := make(jet.VarMap)
	vars.Set("foo", myValue)

	if err := h.App.Render.JetPage(w, r, "sessions", vars, nil); err != nil {
		h.printError("error rendering:", err)
	}

}

func (h *Handlers) JSON(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID      int64    `json:"id"`
		Name    string   `json:"name"`
		Hobbies []string `json:"hobbies"`
	}

	payload.ID = 10
	payload.Name = "pet trak"
	payload.Hobbies = []string{"karate", "tennis", "programming"}

	err := h.App.WriteJson(w, http.StatusOK, payload)
	if err != nil {
		h.printError("cannot write to JSON", err)
	}
}

func (h *Handlers) XML(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ID      int64    `xml:"id"`
		Name    string   `xml:"name"`
		Hobbies []string `xml:"hobbies>hobby"`
	}

	payload.ID = 10
	payload.Name = "pet trak"
	payload.Hobbies = []string{"karate", "tennis", "programming"}

	err := h.App.WriteXML(w, http.StatusOK, payload)
	if err != nil {
		h.printError("cannot write to JSON", err)
	}
}

func (h *Handlers) DownloadFile(w http.ResponseWriter, r *http.Request) {
	h.App.DownloadFile(w, r, "./public/images/", "sokudo.jpg")
}
