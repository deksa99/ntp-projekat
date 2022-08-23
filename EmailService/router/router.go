package router

import (
	"EmailService/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api/email").Subrouter()

	router.HandleFunc("", handler.SendMail).Methods("POST")

	log.Fatal(http.ListenAndServe(":8086", router))
}
