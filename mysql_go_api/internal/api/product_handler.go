package api

import (
	"net/http"

	"github.com/Utsavch189/api_mysql/internal/middlewares"
	"github.com/Utsavch189/api_mysql/internal/services"
	"github.com/gorilla/mux"
)

func ProductHandler(r *mux.Router) {
	r.HandleFunc("/api/products", services.GetAllProducts).
		Methods("GET")
	r.HandleFunc("/api/product/{id}", services.GetAProduct).
		Methods("GET")
	r.Handle("/api/product", middlewares.Authorization(http.HandlerFunc(services.CreateProduct))).
		Methods("POST")
	r.Handle("/api/product", middlewares.Authorization(http.HandlerFunc(services.UpdateAProduct))).
		Methods("PUT")
	r.Handle("/api/product/{id}", middlewares.Authorization(http.HandlerFunc(services.DeleteAProduct))).
		Methods("Delete")
}
