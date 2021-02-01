package api

import "net/http"

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	responseError(w, http.StatusBadRequest, "demo bad request response!")
}
