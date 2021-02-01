package api

import (
	"net/http"
	"vercel-api-demo/app"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	app.ResponseSuccess(w, "Hello world!")
}
