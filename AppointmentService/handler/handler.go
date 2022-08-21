package handler

import (
	"AppointmentService/model"
	"AppointmentService/request"
	"AppointmentService/response"
	"AppointmentService/service"
	"AppointmentService/util/api"
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

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	var acc request.AcceptAppointmentRequest

	err := helper.DecodeJSONBody(w, r, &acc)

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
	rId, _ := strconv.ParseUint(params["requestId"], 10, 32)
	wId, _ := strconv.ParseUint(params["workerId"], 10, 32)

	appointment, err := service.AcceptRequest(uint(rId), uint(wId), acc.Start)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(appointment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Email notification
		errCh := make(chan error)
		go api.RequestAccepted(&appointment, errCh)
	}
}

func RejectRequest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rId, _ := strconv.ParseUint(params["requestId"], 10, 32)
	wId, _ := strconv.ParseUint(params["workerId"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	req, err := service.RejectRequest(uint(rId), uint(wId))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Email notification
		errCh := make(chan error)
		go api.RequestRejected(&req, errCh)
	}
}

func FinishAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	aId, _ := strconv.ParseUint(params["appointmentId"], 10, 32)
	wId, _ := strconv.ParseUint(params["workerId"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	appointment, err := service.FinishCancelAppointment(uint(aId), uint(wId), model.Finished)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(appointment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Email notification
		errCh := make(chan error)
		go api.AppointmentFinished(&appointment, errCh)
	}
}

func CancelAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	aId, _ := strconv.ParseUint(params["appointmentId"], 10, 32)
	wId, _ := strconv.ParseUint(params["workerId"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	appointment, err := service.FinishCancelAppointment(uint(aId), uint(wId), model.CancelledAppointment)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(appointment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Email notification
		errCh := make(chan error)
		go api.AppointmentCancelled(&appointment, errCh)
	}
}

func GetRequestsForCarService(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	carServiceId, _ := strconv.ParseUint(params["carServiceId"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	requests, err := service.GetRequestsForCarService(uint(carServiceId))

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

func GetAppointmentsForWorker(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	workerId, _ := strconv.ParseUint(params["workerId"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	appointments, err := service.GetAppointmentsForWorker(uint(workerId))

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

func GetAppointmentForId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	app, err := service.GetAppointmentById(uint(id))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(app)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
