package main

import ( // Import the missing package

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jcanonbenavi/code/scaff/handler"
)

func main() {
	r := chi.NewRouter()
	h := handler.NewHandler()
	r.Get("/", h.Get())
	r.Get("/users/{userId}", h.GetById())
	r.Get("/users", h.GetByQuery())
	http.ListenAndServe(":8080", r)

}
