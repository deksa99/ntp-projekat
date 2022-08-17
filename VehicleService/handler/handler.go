package handler

import (
	"VehicleService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetVehiclesForUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["userId"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	vehicles, err := service.FindVehiclesForUser(uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(vehicles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func AddVehicle(w http.ResponseWriter, r *http.Request) {

}

func UpdateVehicle(w http.ResponseWriter, r *http.Request) {

}
