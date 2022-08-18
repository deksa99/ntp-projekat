package router

import (
	"VehicleService/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api/vehicles").Subrouter()

	router.HandleFunc("/{vehicleId}", handler.GetVehicle).Methods("GET")
	router.HandleFunc("/all/{userId}", handler.GetVehiclesForUser).Methods("GET")
	router.HandleFunc("/add/{userId}", handler.AddVehicle).Methods("POST")
	router.HandleFunc("/{userId}/{vehicleId}", handler.UpdateVehicle).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8082", router))
}
