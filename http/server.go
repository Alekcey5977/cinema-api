package http

import "github.com/gorilla/mux"

type HTTPServer struct {
	httpHandles *HTTPHandlers
}

func NewHTTPServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandles: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/movie").Methods("Post").HandlerFunc(s.httpHandles.HandlerCteateMovie)
	router.Path("/movie/{title}").Methods("GET").HandlerFunc(s.httpHandles.HandlerGetMovie)
	router.Path("/movie").Methods("GET").HandlerFunc(s.httpHandles.HandlerGetAllMovie)
	router.Path("/movie").Methods("PATCH").HandlerFunc(s.httpHandles.HandlerChangeRating)
	router.Path("/movie").Methods("GET").Queries("adult", "true").HandlerFunc(s.httpHandles.HandlerGetAdultMovie)
	router.Path("/movie").Methods("GET").Queries("adult", "false").HandlerFunc(s.httpHandles.HandlerGetNotAdultMovie)
	router.Path("/movie").Methods("DELETE").HandlerFunc(s.httpHandles.HandlerDeleteMovie)
	
	return nil
}