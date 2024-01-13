package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Data struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type BodyRequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

//var products []Data

type ProductsData struct {
	products []Data
}

func (p *ProductsData) LoadDataFromJson() (product ProductsData, err error) {
	fileContent, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileContent, &p.products)
	fmt.Printf("%T\n", p.products)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (p *ProductsData) getProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p.LoadDataFromJson()
		if err := json.NewEncoder(w).Encode(&p.products); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("wrong request"))
			return
		}
		code := http.StatusOK
		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")

	}
}

func (p *ProductsData) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		for _, product := range p.products {
			if product.Id == id {
				w.WriteHeader(http.StatusCreated)
				w.Header().Set("Content-Type", "application/json")
				responseJSON := BodyRequest(product)
				err := json.NewEncoder(w).Encode(responseJSON)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

func main() {
	router := chi.NewRouter()
	productsdata := ProductsData{}
	router.Route("/", func(r chi.Router) {
		router.Get("/products", productsdata.getProducts())
		// user by id
		router.Get("/products/{id}", productsdata.GetById())

	})
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
