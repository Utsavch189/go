package api

import (
	"github.com/Utsavch189/api_mysql/internal/services"
	"github.com/gorilla/mux"
)

func AuthHandler(r *mux.Router) {
	r.HandleFunc("/api/auth/register", services.UserRegister).Methods("POST")
	r.HandleFunc("/api/auth/login", services.UserLogin).Methods("POST")
}
