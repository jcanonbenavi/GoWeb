package internal

import "errors"

var (
	ErrFieldRequired        = errors.New("Field required")
	ErrFieldQuality         = errors.New("Field quality")
	ErrProductAlreadyExists = errors.New("Product already exists")
)

type ProductService interface {
	Save(product *Product) (err error)
	Get() (elements []Product, err error)
}
