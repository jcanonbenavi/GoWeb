package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jcanonbenavi/app/internal"
)

// type DefaultProduct struct {
// 	product map[int]internal.Product
// 	lastID  int
// }

type BodyRequest struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}
type ProductsData struct {
	products []*internal.Product
}

func (p *ProductsData) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body BodyRequestProductJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("wrong request"))
			return
		}
		p.products = append(p.products, &internal.Product{
			Id:          len(p.products) + 1,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		})
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p.products)

	}
}

func (p *ProductsData) LoadDataFromJson() (product ProductsData, err error) {
	fileContent, err := os.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(fileContent, &p.products)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (p *ProductsData) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//p.LoadDataFromJson()
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
				responseJSON := BodyRequest(*product)
				err := json.NewEncoder(w).Encode(responseJSON)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
