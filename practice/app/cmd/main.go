package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jcanonbenavi/app/internal/handlers"
)

// type Data struct {
// 	Id          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	Quantity    int     `json:"quantity"`
// 	CodeValue   string  `json:"code_value"`
// 	IsPublished bool    `json:"is_published"`
// 	Expiration  string  `json:"expiration"`
// 	Price       float64 `json:"price"`
// }

//var products []Data

func main() {
	router := chi.NewRouter()
	productsdata := handlers.ProductsData{}
	router.Route("/", func(r chi.Router) {
		productsdata.LoadDataFromJson()
		router.Get("/products", productsdata.GetProducts())
		// user by id
		router.Get("/products/{id}", productsdata.GetById())

		router.Post("/products", productsdata.Create())

	})
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
