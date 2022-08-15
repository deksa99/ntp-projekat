package rounter

import (
	"UserService/handler"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/users/register", handler.Register).Methods("POST")
	router.HandleFunc("/api/users/change-password", handler.ChangePassword).Methods("POST")
	router.HandleFunc("/api/users/{id}", handler.FindUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", handler.UpdateUser).Methods("PATCH")
	router.HandleFunc("/api/users", handler.Login).Methods("GET")
	router.HandleFunc("/api/users/{id}/block", handler.BlockUser).Methods("PATCH")
}
