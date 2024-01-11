package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jcanonbenavi/code/internal"
)

// This type represents a collection of movies (movies) stored as a map with integer keys and internal.Movie values.
// It also has a field lastID to keep track of the last assigned ID
type DefaultMovies struct {
	movies map[int]internal.Movie
	lastID int
}

//This function is a "constructor" for creating a new instance of Movie structure.
//It takes a map of movies and the last ID as parameters, and returns a pointer to a DefaultMovies object.

func NewDefaultMovies(movies map[int]internal.Movie, lastId int) *DefaultMovies {
	return &DefaultMovies{
		movies: movies,
		lastID: lastId,
	}
}

type MovieJSON struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Rated     string `json:"rated"`
	Published bool   `json:"published"`
}

// This type represents the expected structure of the JSON body for creating a new movie.
// It includes fields for Title, Year, Rated, and Published.
type BodyRequestMovieJSON struct {
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Rated     string `json:"rated"`
	Published bool   `json:"published"`
}

func (d *DefaultMovies) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// It decodes the JSON body of the request into a BodyRequestMovieJSON struct.
		// If the decoding fails, it responds with a 400 Bad Request status and an "invalid body" message.
		// Otherwise, it creates a new internal.Movie object from the decoded data, increments the lastID,
		//assigns the new ID to the movie, and adds it to the map of movies.
		var body BodyRequestMovieJSON
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("invalid body"))
			return
		}

		movie := internal.Movie{
			Title:     body.Title,
			Year:      body.Year,
			Rated:     body.Rated,
			Published: body.Published,
		}

		(*d).lastID++
		movie.ID = (*d).lastID

		// - validate business rules
		if err := ValidateBussinessRule(&movie); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid movie"))
			return
		}

		(*d).movies[movie.ID] = movie

		//serialize movieJson
		movieJson := MovieJSON{
			ID:        movie.ID,
			Title:     movie.Title,
			Year:      movie.Year,
			Rated:     movie.Rated,
			Published: movie.Published,
		}
		//headers
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		//to bytes
		response := map[string]any{
			"message": "movie created successfully",
			"movie":   movieJson,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func ValidateBussinessRule(mv *internal.Movie) error {
	// validate movie (business rules)
	// - required
	if mv.Title == "" {
		return errors.New("title is required")
	}
	if mv.Year == 0 {
		return errors.New("year is required")
	}
	// - quality
	if len(mv.Title) < 3 || len(mv.Title) > 150 {
		return errors.New("title must be between 3 and 150 characters")
	}
	if mv.Year < 1888 || mv.Year > 2021 {
		return errors.New("year must be between 1888 and 2021")
	}
	// regex ...

	return nil
}
