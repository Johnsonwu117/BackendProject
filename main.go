package main

import (
	"backendProject/routers"
	"net/http"
)

func main() {
	router := routers.Router()

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	s.ListenAndServe()
}
