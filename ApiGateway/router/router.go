package router

import (
	"ApiGateway/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	//User service
	router.Path("/users/login").Queries("role", "{role:user|admin|worker}").HandlerFunc(handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/register", handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/change-password", handler.ChangePassword).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/users/{id}", handler.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/users", handler.GetAllUsers).Methods("GET", "OPTIONS")
	router.HandleFunc("/users/{id}/block", handler.BlockUser).Methods("PATCH", "OPTIONS")

	//Vehicle service
	router.HandleFunc("/vehicles", handler.AddVehicle).Methods("POST", "OPTIONS")
	router.HandleFunc("/vehicles/{vehicleId}", handler.UpdateVehicle).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/vehicles", handler.GetVehiclesForUser).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8090", router))
}
