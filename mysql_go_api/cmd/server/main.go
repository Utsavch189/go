package main

import (
	"net/http"

	"github.com/Utsavch189/api_mysql/internal/api"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api.ProductHandler(r)
	api.AuthHandler(r)
	http.ListenAndServe(":8000", r)
}
