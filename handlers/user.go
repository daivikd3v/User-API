package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/daivikd3v/User-API/data"
	"github.com/daivikd3v/User-API/util"
	guuid "github.com/google/uuid"
	"gopkg.in/go-playground/validator.v9"
)

//User handler Struct
type User struct {
}

//GetUserHandler creates and returns an  instance of User Handler
func GetUserHandler() *User {
	return &User{}
}

//Post creates a users in memory from the request.
func (user User) post(w http.ResponseWriter, r *http.Request) {

	u, err := unmarshal(r)

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validate(u)

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	u.Create()

	util.RespondWithStatus(w, http.StatusOK, true, u)
}

//Put updates an already existing user in memory from the request.
func (user User) put(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	u, err := unmarshal(r)

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = validate(u)

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	u.Uuid, err = guuid.Parse(vars["uuid"])

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid UUID")
		return
	}

	err = u.Update()

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondWithStatus(w, http.StatusOK, true, u)
}

//Delete deletes a user from UUID in the request
func (user User) delete(w http.ResponseWriter, r *http.Request) {

	u := data.User{}

	vars := mux.Vars(r)

	Uuid, err := guuid.Parse(vars["uuid"])

	u.Uuid = Uuid

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid UUID")
		return
	}

	err = u.Delete()

	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	util.RespondWithStatus(w, http.StatusOK, true, u)
}

// Decode json in user struct
func unmarshal(r *http.Request) (*data.User, error) {
	var u data.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		return nil, err
	}

	defer r.Body.Close()

	return &u, nil
}

//Validate request parameters and return error
func validate(u *data.User) error {
	validate := validator.New()
	err := validate.Struct(*u)

	return err
}
