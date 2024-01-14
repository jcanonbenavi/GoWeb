package repository

import "github.com/jcanonbenavi/code/internal"

type MovieMap struct {
	db     map[int]internal.Movie
	lastId int
}

func (m *MovieMap) Save(movie *internal.Movie) (err error) {
	for _, v := range (*m).db {
		if v.ID == movie.ID {
			return internal.ErrMovieAlreadyExists
		}
	}
	(*m).lastId++
	movie.ID = (*m).lastId
	(*m).db[movie.ID] = *movie
	return

}
