package handlers

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func HandleSwagger(w http.ResponseWriter, r *http.Request) {
	httpSwagger.WrapHandler(w, r)
}
