package http

import (
	"cinema-api/cinema"
	"net/http"
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
	
}

func HandlerGetMovie(w http.ResponseWriter, r *http.Request) {

}

func HandlerGetAllMovie(w http.ResponseWriter, r *http.Request) {

}

func HandlerChangeRating(w http.ResponseWriter, r *http.Request) {

}

func HandlerDeleteMovie(w http.ResponseWriter, r *http.Request) {

}