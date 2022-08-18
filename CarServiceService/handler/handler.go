package handler

import (
	"CarServiceService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	user, err := service.GetServicesForCarService(uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
