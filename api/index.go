package api

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	responseSuccess(w, "Hello world!")
}
