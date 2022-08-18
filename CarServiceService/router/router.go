package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api/car-services").Subrouter()

	log.Fatal(http.ListenAndServe(":8083", router))
}
