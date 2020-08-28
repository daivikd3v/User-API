//Package main is the entry point of the API. It takes care of setting up and starting the server.
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/daivikd3v/User-API/handlers"
	muxhandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	handlers.RegisterRoutes(router)

	originsOk := muxhandlers.AllowedOrigins([]string{"*"})
	headersOk := muxhandlers.AllowedHeaders([]string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"})
	methodsOk := muxhandlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions})

	s := &http.Server{
		Addr:         ":8080",
		Handler:      muxhandlers.CORS(originsOk, headersOk, methodsOk)(router),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
