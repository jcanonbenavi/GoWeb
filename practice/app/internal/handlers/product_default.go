package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jcanonbenavi/app/internal"
	"github.com/jcanonbenavi/app/platform/web/request"
	"github.com/jcanonbenavi/app/platform/web/response"
)

type DefaultProduct struct {
	service internal.ProductService
}

func NewDefaultProduct(service internal.ProductService) *DefaultProduct {
	return &DefaultProduct{
		service: service,
	}
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

type BodyRequestProductJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (p *DefaultProduct) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var body BodyRequestProductJSON
		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "Invalid request body")
		}

		products := internal.Product{
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}
		if err := p.service.Save(&products); err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				response.Text(w, http.StatusBadRequest, "invalid body")
			case errors.Is(err, internal.ErrProductAlreadyExists):
				response.Text(w, http.StatusConflict, "Product already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return

		}

		data := BodyRequest{
			Id:          products.Id,
			Name:        products.Name,
			Quantity:    products.Quantity,
			CodeValue:   products.CodeValue,
			IsPublished: products.IsPublished,
			Expiration:  products.Expiration,
			Price:       products.Price,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Product created successfully",
			"data":    data,
		})
	}
}

func (p *DefaultProduct) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		elements, err := p.service.Get()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, v := range elements {
			products := BodyRequest{
				Id:          v.Id,
				Name:        v.Name,
				Quantity:    v.Quantity,
				CodeValue:   v.CodeValue,
				IsPublished: v.IsPublished,
				Expiration:  v.Expiration,
				Price:       v.Price,
			}
			response.JSON(w, http.StatusOK, map[string]any{
				"data": products,
			})
		}
	}

}

// func (p *ProductsData) GetById() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
// 		for _, product := range p.products {
// 			if product.Id == id {
// 				w.WriteHeader(http.StatusCreated)
// 				w.Header().Set("Content-Type", "application/json")
// 				responseJSON := BodyRequest(*product)
// 				err := json.NewEncoder(w).Encode(responseJSON)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 			}
// 		}
// 	}
// }
