//Package handlers implements methods to serve http requests for the API.
package handlers

import (
	"net/http"
)

//RegisterRoutes registers routes to the appropriate handler functions on the ServeMux instance.
func RegisterRoutes(handler *http.ServeMux) {

	homeHandler := GetHomeHandler()
	userHandler := GetUserHandler()

	handler.HandleFunc("/get/", homeHandler.get)
	handler.HandleFunc("/post/", userHandler.post)
	handler.HandleFunc("/put/", userHandler.put)
	handler.HandleFunc("/delete/", userHandler.delete)
	handler.HandleFunc("/", homeHandler.index)

}
