package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jcanonbenavi/code/internal"
	"github.com/jcanonbenavi/code/internal/handler"
	"github.com/jcanonbenavi/code/internal/repository"
	"github.com/jcanonbenavi/code/internal/service"
)

func NewDefaultHTTP(addr string) *DefaultHHTTP {
	return &DefaultHHTTP{
		addr: addr,
	}
}

type DefaultHHTTP struct {
	addr string
}

func (d *DefaultHHTTP) Run() (err error) {
	repository := repository.NewMovieMap(make(map[int]internal.Movie), 0)
	service := service.NewMovieDefault(repository)
	handler := handler.NewMovieDefault(service)

	router := chi.NewRouter()
	router.Route("/movies", func(router chi.Router) {
		router.Get("/{id}", handler.GetByID())
		router.Post("/", handler.Create())
		router.Put("/{id}", handler.Update())
		router.Patch("/{id}", handler.UpdatePartial())
		router.Delete("/{id}", handler.Delete())
	})
	err = http.ListenAndServe(d.addr, router)
	return

}
