package api

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func responseSuccess(w http.ResponseWriter, data interface{}) {
	r := response{Data: data}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}

func responseError(w http.ResponseWriter, code int, message string) {
	r := response{Error: message}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r)
}
