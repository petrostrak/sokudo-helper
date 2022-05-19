package handlers

import "net/http"

func (h *Handlers) ShowCachePage(w http.ResponseWriter, r *http.Request) {
	if err := h.render(w, r, "cache", nil, nil); err != nil {
		h.printError("error rendering:", err)
	}
}

func (h *Handlers) SaveInCache(w http.ResponseWriter, r *http.Request) {}

func (h *Handlers) GetFromCache(w http.ResponseWriter, r *http.Request) {}

func (h *Handlers) DeleteFromCache(w http.ResponseWriter, r *http.Request) {}

func (h *Handlers) EmptyCache(w http.ResponseWriter, r *http.Request) {}
