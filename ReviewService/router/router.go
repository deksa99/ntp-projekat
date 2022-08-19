package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().PathPrefix("/api/reviews").Subrouter()

	log.Fatal(http.ListenAndServe(":8085", router))
}
