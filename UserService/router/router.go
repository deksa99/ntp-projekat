package router

import (
	"UserService/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api/users").Subrouter()

	router.Path("/login").Queries("role", "{role:user|admin|worker}").HandlerFunc(handler.Login).Methods("POST")
	router.HandleFunc("/register", handler.Register).Methods("POST")
	router.HandleFunc("/change-password", handler.ChangePassword).Methods("PATCH")
	router.HandleFunc("/{id}", handler.FindUser).Methods("GET")
	router.HandleFunc("/{id}", handler.UpdateUser).Methods("PATCH")
	router.HandleFunc("", handler.FindAllUsers).Methods("GET")
	router.HandleFunc("/{id}/block", handler.BlockUser).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8081", router))
}
