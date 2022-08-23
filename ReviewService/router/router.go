package router

import (
	"ReviewService/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api/reviews").Subrouter()

	router.HandleFunc("/reported", handler.GetReportedReviews).Methods("GET")
	router.HandleFunc("/car-service/{carServiceId}", handler.GetReviews).Methods("GET")
	router.HandleFunc("/add/{userId}", handler.CreateReview).Methods("POST")
	router.HandleFunc("/{reviewId}/report", handler.ReportReview).Methods("PATCH")
	router.HandleFunc("/report/{reportId}/process", handler.ProcessReport).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8085", router))
}
