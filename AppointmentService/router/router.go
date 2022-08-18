package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api/appointments").Subrouter()

	log.Fatal(http.ListenAndServe(":8084", router))
}
