package handler

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Photo IPhoto
}

func New() *Handler {
	return &Handler{
		Photo: newPhoto(),
	}
}

type Response struct {
	Code          int
	Message       string
	InternalError error
}

func (h *Handler) handleResponse(w http.ResponseWriter, code int, msg string, err error) {

	resp := &Response{code, msg, err}

	b, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	w.Write(b)
}
