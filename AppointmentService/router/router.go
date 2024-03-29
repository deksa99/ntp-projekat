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
	router.HandleFunc("/user/{userId}", handler.GetAppointmentsForUser).Methods("GET")
	router.HandleFunc("/requests/{requestId}/{workerId}/accept", handler.AcceptRequest).Methods("PATCH")
	router.HandleFunc("/requests/{requestId}/{workerId}/reject", handler.RejectRequest).Methods("PATCH")
	router.HandleFunc("/{appointmentId}/{workerId}/finish", handler.FinishAppointment).Methods("PATCH")
	router.HandleFunc("/{appointmentId}/{workerId}/cancel", handler.CancelAppointment).Methods("PATCH")
	router.HandleFunc("/requests/service/{carServiceId}", handler.GetRequestsForCarService).Methods("GET")
	router.HandleFunc("/car-service/{carServiceId}", handler.GetAppointmentsForCarService).Methods("GET")
	router.HandleFunc("/worker/{workerId}", handler.GetAppointmentsForWorker).Methods("GET")
	router.HandleFunc("/{id}", handler.GetAppointmentForId).Methods("GET")

	log.Fatal(http.ListenAndServe(":8084", router))
}
