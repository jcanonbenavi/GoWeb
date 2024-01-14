package internal

import "errors"

var (
	ErrFieldRequired        = errors.New("field required")
	ErrFieldQuality         = errors.New("field quality")
	ErrorMovieAlreadyExists = errors.New("movie already exists")
)

type MovieService interface {
	Save(movie *Movie) (err error)
}
