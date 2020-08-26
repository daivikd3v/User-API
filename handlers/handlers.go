//Package handlers implements methods to serve http requests for the API.
package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

//RegisterRoutes registers routes to the appropriate handler functions on the Mux Router instance.
func RegisterRoutes(handler *mux.Router) {

	homeHandler := GetHomeHandler()
	userHandler := GetUserHandler()

	handler.HandleFunc("/get/", homeHandler.get).Methods(http.MethodGet)
	handler.HandleFunc("/post/", userHandler.post).Methods(http.MethodPost)
	handler.HandleFunc("/put/{uuid}", userHandler.put).Methods(http.MethodPut)
	handler.HandleFunc("/delete/{uuid}", userHandler.delete).Methods(http.MethodDelete)
	handler.HandleFunc("/", homeHandler.index).Methods(http.MethodGet)

}
