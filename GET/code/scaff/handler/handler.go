package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MyHandler struct {
	data map[string]string
}

func NewHandler() *MyHandler {
	return &MyHandler{
		data: map[string]string{
			"1": "Juan",
			"2": "Pedro",
			"3": "Maria",
			"4": "Jose",
		},
	}
}

func (h *MyHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	}
}

type MyResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (h *MyHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userId")
		name, ok := h.data[id]
		if !ok {
			code := http.StatusNotFound
			body := MyResponse{
				Message: "user not found",
				Data:    nil,
			}
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			return
		}

		code := http.StatusOK
		body := MyResponse{Message: "user found", Data: name}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
		return
	}
}

func (h *MyHandler) GetByQuery() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("userId")
		hobby := r.URL.Query().Get("hobby")

		name, ok := h.data[id]
		if !ok {
			code := http.StatusNotFound
			body := MyResponse{Message: "user not found", Data: nil}
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			return
		}
		code := http.StatusOK
		body := MyResponse{Message: "user found", Data: name + " " + hobby}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)

	}
}
