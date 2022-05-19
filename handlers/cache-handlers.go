package handlers

import "net/http"

func (h *Handlers) ShowCachePage(w http.ResponseWriter, r *http.Request) {
	if err := h.render(w, r, "cache", nil, nil); err != nil {
		h.printError("error rendering:", err)
	}
}

func (h *Handlers) SaveInCache(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Name  string `json:"name"`
		Value string `json:"value"`
		CSRF  string `json:"csrf_token"`
	}

	err := h.App.ReadJSON(w, r, &userInput)
	if err != nil {
		h.App.Error500(w, r)
		return
	}

	err = h.App.Cache.Set(userInput.Name, userInput.Value)
	if err != nil {
		h.App.Error500(w, r)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "Saved in cache"

	_ = h.App.WriteJson(w, http.StatusCreated, resp)
}

func (h *Handlers) GetFromCache(w http.ResponseWriter, r *http.Request) {
	var msg string
	var inCache = true

	var userInput struct {
		Name string `json:"name"`
		CSRF string `json:"csrf_token"`
	}

	err := h.App.ReadJSON(w, r, &userInput)
	if err != nil {
		h.App.Error500(w, r)
		return
	}

	fromCache, err := h.App.Cache.Get(userInput.Name)
	if err != nil {
		msg = "Not found in cache!"
		inCache = false
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		Value   string `json:"value"`
	}

	if inCache {
		resp.Error = false
		resp.Message = "success"
		resp.Value = fromCache.(string)
	} else {
		resp.Error = true
		resp.Message = msg
	}

	_ = h.App.WriteJson(w, http.StatusCreated, resp)
}

func (h *Handlers) DeleteFromCache(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Name string `json:"name"`
		CSRF string `json:"csrf_token"`
	}

	err := h.App.ReadJSON(w, r, &userInput)
	if err != nil {
		h.App.Error500(w, r)
		return
	}

	err = h.App.Cache.Forget(userInput.Name)
	if err != nil {
		h.App.Error500(w, r)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "deleted from cache (if it existed)"

	_ = h.App.WriteJson(w, http.StatusCreated, resp)
}

func (h *Handlers) EmptyCache(w http.ResponseWriter, r *http.Request) {}
