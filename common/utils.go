package common

import (
	"encoding/json"
	"net/http"
)

func RespondWithStatus(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
