package api

import (
	"net/http"

	"github.com/Utsavch189/api_mysql/internal/middlewares"
	"github.com/Utsavch189/api_mysql/internal/services"
	"github.com/gorilla/mux"
)

func ProductHandler(r *mux.Router) {
	r.Handle("/api/products", middlewares.Authorization(http.HandlerFunc(services.GetAllProducts))).
		Methods("GET")
	r.HandleFunc("/api/product/{id}", services.GetAProduct).Methods("GET")
	r.HandleFunc("/api/product", services.CreateProduct).Methods("POST")
	r.HandleFunc("/api/product", services.UpdateAProduct).Methods("PUT")
	r.HandleFunc("/api/product/{id}", services.DeleteAProduct).Methods("Delete")
}
