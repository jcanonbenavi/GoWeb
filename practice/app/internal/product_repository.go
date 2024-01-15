package internal

import "errors"

var (
	ErrProductNameAlreadyExists = errors.New("Product already exists")
	ErrNotFound                 = errors.New("Product not found")
)

type ProductRepository interface {
	Save(product *Product) (err error)
	Get() (elements []Product, err error)
}
