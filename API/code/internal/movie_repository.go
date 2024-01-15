package internal

import "errors"

var (
	ErrMovieTitleAlreadyExists = errors.New("movie already exists")
	ErrMovieNotFound           = errors.New("movie not found")
)

type MovieRepository interface {
	Save(movie *Movie) (err error)
	GetByID(id int) (movie Movie, err error)
	Update(movie *Movie) (err error)
	Delete(id int) (err error)
}
