package handler

import (
	"ApiGateway/util/auth"
	"ApiGateway/util/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddVehicle(w http.ResponseWriter, r *http.Request) {
	resp, err := auth.Authentication(w, r, []string{"user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.VehicleServiceRoundRobin.Next().Host + "/api/vehicles/add/" + strconv.FormatUint(uint64(resp.UserId), 10)
	response.HandleRequest(w, r, url)
}

func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vId, _ := strconv.ParseUint(params["vehicleId"], 10, 32)

	resp, err := auth.Authentication(w, r, []string{"user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.VehicleServiceRoundRobin.Next().Host + "/api/vehicles/" +
		strconv.FormatUint(uint64(resp.UserId), 10) + "/" + strconv.FormatUint(vId, 10)
	response.HandleRequest(w, r, url)
}

func GetVehiclesForUser(w http.ResponseWriter, r *http.Request) {
	resp, err := auth.Authentication(w, r, []string{"user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.VehicleServiceRoundRobin.Next().Host + "/api/vehicles/all/" + strconv.FormatUint(uint64(resp.UserId), 10)
	response.HandleRequest(w, r, url)
}
