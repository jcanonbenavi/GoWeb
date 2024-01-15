package internal

import (
	"errors"
)

var (
	ErrFieldRequired      = errors.New("field required")
	ErrFieldQuality       = errors.New("field quality")
	ErrMovieAlreadyExists = errors.New("movie already exists")
)

type MovieService interface {
	Save(movie *Movie) (err error)
	GetByID(id int) (movie Movie, err error)
	Update(movie *Movie) (err error)
	Delete(id int) (err error)
}
