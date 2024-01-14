package service

import (
	"fmt"

	"github.com/jcanonbenavi/code/internal"
)

func NewMovieDefault(rp internal.MovieRepository) *MovieDefault {
	return &MovieDefault{
		rp: rp,
	}
}

type MovieDefault struct {
	//repository
	rp internal.MovieRepository
}

func (m *MovieDefault) Save(movie *internal.Movie) (err error) {
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

	err = m.rp.Save(movie)
	if err != nil {
		switch err {
		case internal.ErrMovieAlreadyExists: //error return by repository
			return fmt.Errorf("%w: movie", internal.ErrorMovieAlreadyExists) //service error
		}
		return
	}
	return
}
