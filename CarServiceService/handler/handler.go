package handler

import (
	"CarServiceService/request"
	"CarServiceService/response"
	"CarServiceService/service"
	"CarServiceService/util/converter"
	"CarServiceService/util/helper"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
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
	var myLocation request.Location

	err := helper.DecodeJSONBody(w, r, &myLocation)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		var e *response.Error
		if errors.As(err, &e) {
			http.Error(w, e.Message, e.Status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	carService, err := service.FindNearestService(myLocation.Latitude, myLocation.Longitude)

	if err != nil {
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(converter.CarServiceToCarServiceInfo(&carService))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func CreateService(w http.ResponseWriter, r *http.Request) {

}

func UpdateService(w http.ResponseWriter, r *http.Request) {

}

func ChangeAvailability(w http.ResponseWriter, r *http.Request) {

}

func CreateCarService(w http.ResponseWriter, r *http.Request) {

}
