package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jcanonbenavi/app/internal/handlers"
	"github.com/jcanonbenavi/app/internal/loader"
	"github.com/jcanonbenavi/app/internal/repository"
	"github.com/jcanonbenavi/app/internal/service"
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
	products, err := loader.LoadDataFromJson()
	repository := repository.NewProductSlice(products)
	//repository := repository.NewProductSlice([]internal.Product{})
	service := service.NewProductDefault(repository)
	handlers := handlers.NewDefaultProduct(service)
	router := chi.NewRouter()
	router.Post("/products", handlers.Create())
	router.Get("/products", handlers.Get())
	err = http.ListenAndServe(d.addr, router)
	return
}
