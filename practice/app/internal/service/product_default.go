package service

import (
	"fmt"

	"github.com/jcanonbenavi/app/internal"
)

type ProductDefault struct {
	repository internal.ProductRepository
}

func NewProductDefault(repository internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		repository: repository,
	}
}

func (p *ProductDefault) Save(product *internal.Product) (err error) {
	if (*product).Expiration == "" {
		return fmt.Errorf("%w: Expiration", internal.ErrFieldRequired)
	}
	if (*product).Name == "" {
		return fmt.Errorf("%w: Name", internal.ErrFieldRequired)
	}
	if (*product).Price == 0 {
		return fmt.Errorf("%w: Price", internal.ErrFieldRequired)
	}
	if (*product).CodeValue == "" {
		return fmt.Errorf("%w: Code Value", internal.ErrFieldRequired)
	}
	if (*product).Quantity == 0 {
		return fmt.Errorf("%w: Quantity", internal.ErrFieldRequired)
	}
	err = p.repository.Save(product)
	if err != nil {
		switch err {
		case internal.ErrProductNameAlreadyExists:
			return fmt.Errorf("%w: Name", internal.ErrProductAlreadyExists)
		}
		return

	}
	return
}

func (p *ProductDefault) Get() (elements []internal.Product, err error) {
	return p.repository.Get()
}
