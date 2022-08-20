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
	router.HandleFunc("", handler.CreateReview).Methods("POST")
	router.HandleFunc("/{reviewId}/report", handler.ReportReview).Methods("PATCH")
	router.HandleFunc("/{reviewId}/process", handler.ReportReview).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8085", router))
}
