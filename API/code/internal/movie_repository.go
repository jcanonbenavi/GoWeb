package internal

import "errors"

var (
	ErrMovieAlreadyExists = errors.New("movie already exists")
)

type MovieRepository interface {
	Save(movie *Movie) (err error)
}
