package handler

import (
	"AppointmentService/request"
	"AppointmentService/response"
	"AppointmentService/service"
	"AppointmentService/util/helper"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func CreateAppointmentRequest(w http.ResponseWriter, r *http.Request) {
	var cr request.CreateAppointmentRequest

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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		err = json.NewEncoder(w).Encode(newRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func CancelRequest(w http.ResponseWriter, r *http.Request) {
	var cr request.CancelAppointmentRequest

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

	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["userId"], 10, 32)

	if uint(userId) != cr.UserID {
		http.Error(w, "oops you can only cancel your appointment", http.StatusBadRequest)
		return
	}

	updatedRequest, err := service.CancelAppointmentRequest(cr.UserID, cr.AppointmentRequestID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(updatedRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func GetRequestsForUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["userId"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	requests, err := service.GetRequestsForUser(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(requests)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func GetAppointmentsForUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["userId"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	appointments, err := service.GetAppointmentsForUser(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(appointments)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func CreateAppointment(w http.ResponseWriter, r *http.Request) {

}

func ShowNewRequestsForService(w http.ResponseWriter, r *http.Request) {

}

func ShowAppointmentsForWorker(w http.ResponseWriter, r *http.Request) {

}
