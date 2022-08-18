package handler

import (
	"CarServiceService/service"
	"encoding/json"
	"net/http"
)

func GetAllCarServices(w http.ResponseWriter, r *http.Request) {
	services := service.FindAllCarServices()

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(services)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func GetCatalog(w http.ResponseWriter, r *http.Request) {

}

func FindNearest(w http.ResponseWriter, r *http.Request) {

}

func CreateService(w http.ResponseWriter, r *http.Request) {

}

func UpdateService(w http.ResponseWriter, r *http.Request) {

}

func ChangeAvailability(w http.ResponseWriter, r *http.Request) {

}

func CreateCarService(w http.ResponseWriter, r *http.Request) {

}
