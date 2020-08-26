package util

import (
	"encoding/json"
	"net/http"
)

//RespondWithError is a utility function to encode error messages
//    in JSON Response format and writes it to the ResponseWriter.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]interface{}{"success": false, "error": message})
}

//RespondWithStatus is a utility function which encodes data
// in appropriate JSON Response format to the Response Writer
func RespondWithStatus(w http.ResponseWriter, code int, status bool, payload interface{}) {
	RespondWithJSON(w, code, map[string]interface{}{"success": status, "data": payload})
}

//RespondWithJSON encodes Json response and writes it to response writer
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
