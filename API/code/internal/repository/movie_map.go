package repository

import "github.com/jcanonbenavi/code/internal"

type MovieMap struct {
	db     map[int]internal.Movie
	lastId int
}

func NewMovieMap(db map[int]internal.Movie, lastId int) *MovieMap {
	// default config / values
	// ...

	return &MovieMap{
		db:     db,
		lastId: lastId,
	}
}

func (m *MovieMap) Save(movie *internal.Movie) (err error) {
	for _, movies := range (*m).db {
		if movies.Title == movie.Title {
			return internal.ErrMovieTitleAlreadyExists
		}
	}
	(*m).lastId++
	movie.ID = (*m).lastId
	(*m).db[movie.ID] = *movie
	return

}

func (m *MovieMap) GetByID(id int) (movie internal.Movie, err error) {
	movie, ok := m.db[id]
	if !ok {
		err = internal.ErrMovieNotFound
		return
	}

	return
}

func (m *MovieMap) Update(movie *internal.Movie) (err error) {

	_, ok := m.db[(*movie).ID]
	if !ok {
		err = internal.ErrMovieNotFound
		return
	}

	m.db[(*movie).ID] = *movie
	return
}

func (m *MovieMap) Delete(id int) (err error) {
	_, ok := m.db[id]
	if !ok {
		err = internal.ErrMovieNotFound
		return
	}

	delete(m.db, id)

	return
}
