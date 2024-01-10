package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ControllerEmployee struct {
	st map[string]string
}

type ResponseGetByIdEmployee struct {
	Message string `json:"message"`
	Data    *struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
	Error bool `json:"error"`
}

//request

func (c *ControllerEmployee) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		_, ok := c.st[id]

		if !ok {
			code := http.StatusNotFound
			body := &ResponseGetByIdEmployee{Message: "employee not found", Data: nil, Error: true}
			w.WriteHeader(code)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			return
		}
	}
}

//response

// func (c *ControllerEmployee) GetById() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		code := http.StatusOK
// 		body := &ResponseGetByIdEmployee{Message: "employee found", Data: &struct {
// 			Id   string `json:"id"`
// 			Name string `json:"name"`
// 		}{Id: Id, Name: employee}, Error: false},
// 			w.WriteHeader(code)
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(body)
// 	}
// }
