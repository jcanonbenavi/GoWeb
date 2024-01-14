package handler

import (
	"errors"
	"net/http"

	"github.com/jcanonbenavi/code/internal"
	"github.com/jcanonbenavi/code/platform/web/request"
	"github.com/jcanonbenavi/code/platform/web/response"
)

type DefaultMovies struct {
	// sv is a movie service
	sv internal.MovieService
}

func NewMovieDefault(sv internal.MovieService) *DefaultMovies {
	return &DefaultMovies{
		sv: sv,
	}
}

type MovieJSON struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Year      int     `json:"year"`
	Rating    float64 `json:"rating"`
	Published bool    `json:"published"`
}

type BodyRequestMovieJSON struct {
	Title     string  `json:"title"`
	Year      int     `json:"year"`
	Rating    float64 `json:"rating"`
	Published bool    `json:"published"`
}

func (d *DefaultMovies) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var body BodyRequestMovieJSON
		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}
		// process
		// - serialize internal.Movie
		movie := internal.Movie{
			Title:     body.Title,
			Year:      body.Year,
			Rating:    body.Rating,
			Published: body.Published,
		}
		// - save movie
		if err := d.sv.Save(&movie); err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				// w.Header().Set("Content-Type", "text/plain")
				// w.WriteHeader(http.StatusBadRequest)
				// w.Write([]byte("invalid body"))
				response.Text(w, http.StatusBadRequest, "invalid body")
			case errors.Is(err, internal.ErrMovieAlreadyExists):
				response.Text(w, http.StatusConflict, "movie already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize MovieJSON
		data := MovieJSON{
			ID:        movie.ID,
			Title:     movie.Title,
			Year:      movie.Year,
			Rating:    movie.Rating,
			Published: movie.Published,
		}
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusCreated)
		// json.NewEncoder(w).Encode(map[string]any{
		// 	"message": "movie created",
		// 	"data": data,
		// })
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "movie created",
			"data":    data,
		})
	}
}
