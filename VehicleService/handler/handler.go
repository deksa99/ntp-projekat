package handler

import (
	"VehicleService/request"
	"VehicleService/response"
	"VehicleService/service"
	"VehicleService/util/converter"
	"VehicleService/util/helper"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["vehicleId"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	vehicle, err := service.FindVehicleById(uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

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
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["userId"], 10, 32)

	var addVehicleRequest request.AddVehicle

	err := helper.DecodeJSONBody(w, r, &addVehicleRequest)

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

	if uint(id) != addVehicleRequest.UserId {
		http.Error(w, "user id error", http.StatusBadRequest)
		return
	}

	vehicle, err := service.AddVehicle(
		addVehicleRequest.UserId,
		addVehicleRequest.Manufacturer,
		addVehicleRequest.CarModel,
		addVehicleRequest.Color,
		addVehicleRequest.LicencePlate,
		addVehicleRequest.ChassisNumber,
	)

	if err != nil {
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(converter.VehicleToVehicleInfo(&vehicle))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["userId"], 10, 32)
	vehicleId, _ := strconv.ParseUint(params["vehicleId"], 10, 32)

	var updateVehicleRequest request.UpdateVehicle

	err := helper.DecodeJSONBody(w, r, &updateVehicleRequest)

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

	vehicle, err := service.UpdateVehicle(uint(userId), uint(vehicleId), updateVehicleRequest.Color)

	if err != nil {
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(converter.VehicleToVehicleInfo(&vehicle))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
