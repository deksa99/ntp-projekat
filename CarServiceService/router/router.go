package router

import (
	"CarServiceService/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api/car-services").Subrouter()

	router.HandleFunc("", handler.GetAllCarServices).Methods("GET")
	router.HandleFunc("", handler.CreateCarService).Methods("POST")
	router.HandleFunc("/near-me", handler.FindNearest).Methods("POST")
	router.HandleFunc("/service", handler.CreateService).Methods("POST")
	router.HandleFunc("/service", handler.UpdateService).Methods("PATCH")
	router.HandleFunc("/service/{id}/change-availability", handler.ChangeAvailability).Methods("PATCH")
	router.HandleFunc("/{id}/catalog", handler.GetCatalog).Methods("GET")

	log.Fatal(http.ListenAndServe(":8083", router))
}
