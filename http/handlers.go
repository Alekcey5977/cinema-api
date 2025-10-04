package http

import (
	"cinema-api/cinema"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


type HTTPHandlers struct {
	listMovie *cinema.List 
}

/*
pattern: /movie

method: POST

info: JSON in HTTP requst body

succed:
	- status code: 201 Create
	- response body: JSON represent created body

failed:
	- status code: 400 Bad Request, 409 Conflict, 500 ...
	- response body: JSON with error + time
*/
func (h *HTTPHandlers) HandlerCteateMovie(w http.ResponseWriter, r *http.Request) {
	var movie cinema.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		errDTO := NewErrorDTO(err.Error(), time.Now())

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)

		return
	}

	if err := movie.ValidateMovieData(); err != nil {
		errDTO := NewErrorDTO(err.Error(), time.Now())

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	cinemaMovie := cinema.NewMovie(
		movie.Title, movie.Description,
		movie.Genres,
		movie.YearsOfRelease,
		movie.Rating,
		movie.Country,
		movie.Adult,
	)

	if err := h.listMovie.AddMovie(*cinemaMovie); err != nil {
		errDTO := NewErrorDTO(err.Error(), time.Now())

		if errors.Is(err, cinema.ErrMovieAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(cinemaMovie, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}


/*
pattern: /movie/{title}

method: GET

info: pattern

succed:
	- stutus code: 200 Ok
	- respons body: JSON represent found movie

failed:
	- ststus code : 400 Bad Requst, 404 Not Found, 500 ...
	- response body: JSON with error + time 
*/
func (h *HTTPHandlers) HandlerGetMovie(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	movie, err := h.listMovie.GetMovie(title)
	if err != nil {
		errDTO := NewErrorDTO(err.Error(), time.Now())

		if errors.Is(err, cinema.ErrMovieNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)

		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(movie, "", "    ")
	if err != nil {
		panic(err)
	}
	
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response ", err)
		return
	}
}

/*
pattern: /movie

method: GET

info: -

succed:
	- status code: 200 Ok
	- response body: JSON represent found movies

failed:
	- status code: 400 Bad Request, 500 ...
	- response body: JSON  with error + time
*/
func (h *HTTPHandlers) HandlerGetAllMovie(w http.ResponseWriter, r *http.Request) {
	movies := h.listMovie.ListMovies()

	b, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		panic(err)
	}
	
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
	


}

/*
pattern: /movie/{title}

method: PATCH

info: pattern + JSON in requst body

succed:
	- status code: 200 Ok
	- response body: JSON represent found movie

failed:
	- status code: 400 Bad Request, 404, 409 Conflict,  500 ...
	- response body: JSON  with error + time
*/
func(h *HTTPHandlers) HandlerChangeRating(w http.ResponseWriter, r *http.Request) {
	var ratingDTO RatingChangeDTO
	
	if err := json.NewDecoder(r.Body).Decode(&ratingDTO); err !=nil {
		errDTO := NewErrorDTO(err.Error(), time.Now())

		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)

		return
	}

	title := mux.Vars(r)["title"]

	movie, err := h.listMovie.ChangeRatingMovie(title, ratingDTO.Rating)
	if err != nil {
		errDTO := NewErrorDTO(err.Error(), time.Now())

		if errors.Is(err, cinema.ErrIncorrectRating) ||
			errors.Is(err, cinema.ErrMovieNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(movie, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)

		return
	}
}

/*
pattern: /movie?adult=true

method: GET

info: query params

succed:
	- status code: 200 Ok
	- response body: JSON represent found movies

failed:
	- status code: 400 Bad Request, 500 ...
	- response body: JSON  with error + time
*/
func (h *HTTPHandlers) HandlerGetAdultMovie(w http.ResponseWriter, r *http.Request) {
	adultMovies := h.listMovie.GetAdultMovie()

	b, err := json.MarshalIndent(adultMovies, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		
		return
	}
}

/*
pattern: /movie?adult=false

method: GET

info: query params

succed:
	- status code: 200 Ok
	- response body: JSON represent found movies

failed:
	- status code: 400 Bad Request, 500 ...
	- response body: JSON  with error + time
*/
func (h *HTTPHandlers) HandlerGetNotAdultMovie(w http.ResponseWriter, r *http.Request) {
	notAdultMovies := h.listMovie.GetNotAdultMovie()

	b, err := json.MarshalIndent(notAdultMovies, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		
		return
	}

}

/*
pattern: /movie/{title}

method: DELETE

info: pattern + JSON in requst body

succed:
	- status code: 204 Not Comtent
	- response body: -

failed:
	- status code: 400 Bad Request, 404,  500 ...
	- response body: JSON  with error + time
*/
func (h *HTTPHandlers) HandlerDeleteMovie(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.listMovie.DeleteMovie(title); err != nil {
		errDTO := NewErrorDTO(err.Error(), time.Now())

		if errors.Is(err, cinema.ErrMovieNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}