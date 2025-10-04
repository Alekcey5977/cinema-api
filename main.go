package main

import (
	"cinema-api/cinema"
	"cinema-api/http"
	"fmt"
)

func main() {
	cinemaList := cinema.NewList()

	httpHandlers := http.NewHTTPHandlers(cinemaList)

	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}