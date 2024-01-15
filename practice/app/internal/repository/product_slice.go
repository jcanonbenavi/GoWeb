package repository

import (
	"github.com/jcanonbenavi/app/internal"
)

type ProductsData struct {
	productsSlice []internal.Product
}

func NewProductSlice(product []internal.Product) *ProductsData {
	return &ProductsData{
		productsSlice: product,
	}
}

func (p *ProductsData) Save(product *internal.Product) (err error) {
	for _, products := range (*p).productsSlice {
		if products.Name == product.Name {
			return internal.ErrProductNameAlreadyExists
		}
	}
	product.Id = len(p.productsSlice) + 1
	(*p).productsSlice = append((*p).productsSlice, *product)
	return
}

func (p *ProductsData) Get() (elements []internal.Product, err error) {
	elements = p.productsSlice
	if len(elements) == 0 {
		err = internal.ErrNotFound
	}
	return
}
