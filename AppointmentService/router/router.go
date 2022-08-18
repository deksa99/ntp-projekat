package router

import (
	"AppointmentService/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api/appointments").Subrouter()

	router.HandleFunc("/requests", handler.CreateAppointmentRequest).Methods("POST")
	router.HandleFunc("/requests/{userId}/cancel", handler.CancelRequest).Methods("PATCH")
	router.HandleFunc("/requests/{userId}", handler.GetRequestsForUser).Methods("GET")
	router.HandleFunc("/{userId}", handler.GetAppointmentsForUser).Methods("GET")
	router.HandleFunc("/requests/{requestId}/accept", handler.CreateAppointment).Methods("PATCH")
	router.HandleFunc("/requests/service/{serviceId}", handler.ShowNewRequestsForService).Methods("POST")
	router.HandleFunc("/worker/{workerId}", handler.ShowAppointmentsForWorker).Methods("POST")

	log.Fatal(http.ListenAndServe(":8084", router))
}
