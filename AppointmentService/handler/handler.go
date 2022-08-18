package handler

import (
	"AppointmentService/request"
	"AppointmentService/response"
	"AppointmentService/service"
	"AppointmentService/util/helper"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func CreateAppointmentRequest(w http.ResponseWriter, r *http.Request) {
	var cr request.CreateAppointment

	err := helper.DecodeJSONBody(w, r, &cr)

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

	newRequest, err := service.CreateAppointmentRequest(cr.UserID, cr.VehicleID, cr.ServiceID, cr.CarServiceID)

	if err != nil {
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(newRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func CancelRequest(w http.ResponseWriter, r *http.Request) {

}

func GetRequestsForUser(w http.ResponseWriter, r *http.Request) {

}

func GetAppointmentsForUser(w http.ResponseWriter, r *http.Request) {

}

func CreateAppointment(w http.ResponseWriter, r *http.Request) {

}

func ShowNewRequestsForService(w http.ResponseWriter, r *http.Request) {

}

func ShowAppointmentsForWorker(w http.ResponseWriter, r *http.Request) {

}
