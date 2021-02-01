package api

import (
	"net/http"
	"vercel-api-demo/app"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	app.ResponseError(w, http.StatusBadRequest, "demo bad request response!")
}
