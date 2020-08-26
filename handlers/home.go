package handlers

import (
	"net/http"

	"github.com/daivikd3v/User-API/util"
)

//Home Handler Struct
type Home struct {
}

//GetHomeHandler creates an instance of Home handler Struct.
func GetHomeHandler() *Home {
	return &Home{}
}

//Index - returns a welcome message in JSON format.
func (home Home) index(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to home page route",
	}
	util.RespondWithJSON(w, http.StatusOK, response)
}

//Get - returns a message indicating a get route has reached.
func (home Home) get(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "You have reached the get route",
	}
	util.RespondWithJSON(w, http.StatusOK, response)
}
