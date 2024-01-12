package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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

//var products []Data

type ProductsData struct {
	products []Data
}

func (p *ProductsData) LoadDataFromJson() {
	fileContent, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileContent, &p.products)
	if err != nil {
		log.Fatal(err)
	}
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	productsdata := ProductsData{}
	productsdata.LoadDataFromJson()
	if err := json.NewEncoder(w).Encode(productsdata.products); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("wrong request"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

}

func main() {
	router := chi.NewRouter()
	router.Get("/", getProductsHandler)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
