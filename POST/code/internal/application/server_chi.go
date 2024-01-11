package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jcanonbenavi/code/internal"
	"github.com/jcanonbenavi/code/internal/handlers"
)

type serverChi struct {
	address string
}

func NewServerChi(address string) *serverChi {
	defaultAdress := ":8080"
	if address != "" {
		defaultAdress = address
	}
	return &serverChi{
		address: defaultAdress,
	}
}

func (s *serverChi) Run() error {
	db := make(map[int]internal.Movie, 0)
	lastID := 0
	handler := handlers.NewDefaultMovies(db, lastID)

	//router
	router := chi.NewRouter()
	//routes
	router.Post("/movies", handler.Create())

	return http.ListenAndServe(s.address, router)

}
