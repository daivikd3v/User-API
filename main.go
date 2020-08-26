//Package main is the entry point of the API. It takes care of setting up and starting the server.
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/daivikd3v/User-API/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	handlers.RegisterRoutes(router)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
