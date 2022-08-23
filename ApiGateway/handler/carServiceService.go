package handler

import (
	"ApiGateway/util/auth"
	"ApiGateway/util/response"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllCarServices(w http.ResponseWriter, r *http.Request) {
	url := response.CarServiceServiceRoundRobin.Next().Host + "/api/car-services"
	response.HandleRequest(w, r, url)
}

func CreateCarService(w http.ResponseWriter, r *http.Request) {
	_, err := auth.Authentication(w, r, []string{"admin"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.CarServiceServiceRoundRobin.Next().Host + "/api/car-services"
	response.HandleRequest(w, r, url)
}

func FindNearest(w http.ResponseWriter, r *http.Request) {
	_, err := auth.Authentication(w, r, []string{"user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.CarServiceServiceRoundRobin.Next().Host + "/api/car-services/near-me"
	response.HandleRequest(w, r, url)
}

func CreateService(w http.ResponseWriter, r *http.Request) {
	_, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.CarServiceServiceRoundRobin.Next().Host + "/api/car-services/service"
	response.HandleRequest(w, r, url)
}

func UpdateService(w http.ResponseWriter, r *http.Request) {
	_, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.CarServiceServiceRoundRobin.Next().Host + "/api/car-services/service"
	response.HandleRequest(w, r, url)
}

func ChangeAvailability(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := params["id"]

	_, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.CarServiceServiceRoundRobin.Next().Host + "/api/car-services/service/" + id + "/change-availability"
	response.HandleRequest(w, r, url)
}

func GetCatalog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := params["id"]

	url := response.CarServiceServiceRoundRobin.Next().Host + "/api/car-services/" + id + "/catalog"
	response.HandleRequest(w, r, url)
}
