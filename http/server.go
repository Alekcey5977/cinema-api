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


	return nil
}