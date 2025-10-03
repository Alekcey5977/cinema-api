package http

import (
	"cinema-api/cinema"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
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
		errDTO := cinema.NewErrorDTO(err.Error(), time.Now())

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)

		return
	}

	if err := movie.ValidateMovieData(); err != nil {
		errDTO := cinema.NewErrorDTO(err.Error(), time.Now())

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
		errDTO := cinema.NewErrorDTO(err.Error(), time.Now())

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

}