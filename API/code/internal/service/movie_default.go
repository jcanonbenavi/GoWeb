package service

import (
	"fmt"

	"github.com/jcanonbenavi/code/internal"
)

type MovieDefault struct {
	//repository
	rp internal.MovieRepository
}

func NewMovieDefault(rp internal.MovieRepository) *MovieDefault {
	return &MovieDefault{
		rp: rp,
	}
}

func ValidateMovie(movie *internal.Movie) (err error) {
	// - validate required fields
	if (*movie).Title == "" {
		return fmt.Errorf("%w: title", internal.ErrFieldRequired)
	}
	if (*movie).Year == 0 {
		return fmt.Errorf("%w: year", internal.ErrFieldRequired)
	}
	// - validate quality
	if !((*movie).Rating >= 0 && (*movie).Rating <= 10) {
		return fmt.Errorf("%w: rating", internal.ErrFieldQuality)
	}
	if !((*movie).Year >= 1888 && (*movie).Year <= 9999) {
		return fmt.Errorf("%w: year", internal.ErrFieldQuality)
	}

	return
}

func (m *MovieDefault) Save(movie *internal.Movie) (err error) {
	if err = ValidateMovie(movie); err != nil {
		return
	}

	err = m.rp.Save(movie)
	if err != nil {
		switch err {
		case internal.ErrMovieTitleAlreadyExists: //error return by repository
			err = fmt.Errorf("%w: movie", internal.ErrMovieAlreadyExists) //service error
		}
		return
	}
	return
}

func (m *MovieDefault) GetByID(id int) (movie internal.Movie, err error) {
	// get movie
	movie, err = m.rp.GetByID(id)
	if err != nil {
		switch err {
		case internal.ErrMovieNotFound:
			err = fmt.Errorf("%w: id", internal.ErrMovieNotFound)
		}
		return
	}

	return
}

func (m *MovieDefault) Update(movie *internal.Movie) (err error) {
	// validate
	if err = ValidateMovie(movie); err != nil {
		return
	}

	// update movie
	err = m.rp.Update(movie)
	if err != nil {
		switch err {
		case internal.ErrMovieNotFound:
			err = fmt.Errorf("%w: id", internal.ErrMovieNotFound)
		}
		return
	}
	return
}

func (m *MovieDefault) Delete(id int) (err error) {
	// delete movie
	err = m.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrMovieNotFound:
			err = fmt.Errorf("%w: id", internal.ErrMovieNotFound)
		}
		return
	}
	return
}
