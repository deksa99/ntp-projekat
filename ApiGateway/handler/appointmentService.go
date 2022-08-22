package handler

import (
	"ApiGateway/util/auth"
	"ApiGateway/util/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateAppointmentRequest(w http.ResponseWriter, r *http.Request) {
	_, err := auth.Authentication(w, r, []string{"user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/requests"
	response.HandleRequest(w, r, url)
}

func CancelRequest(w http.ResponseWriter, r *http.Request) {
	user, err := auth.Authentication(w, r, []string{"user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	id := strconv.Itoa(int(user.UserId))
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/requests/" + id + "/cancel"
	response.HandleRequest(w, r, url)
}

func GetRequestsForUser(w http.ResponseWriter, r *http.Request) {
	user, err := auth.Authentication(w, r, []string{"user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	id := strconv.Itoa(int(user.UserId))
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/requests/" + id
	response.HandleRequest(w, r, url)
}

func GetAppointmentsForUser(w http.ResponseWriter, r *http.Request) {
	user, err := auth.Authentication(w, r, []string{"user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	id := strconv.Itoa(int(user.UserId))
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/user" + id
	response.HandleRequest(w, r, url)
}

func AcceptRequest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rId := params["requestId"]

	worker, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	id := strconv.Itoa(int(worker.UserId))
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/requests/" + rId + "/" + id + "/accept"
	response.HandleRequest(w, r, url)
}

func RejectRequest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	rId := params["requestId"]

	worker, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	id := strconv.Itoa(int(worker.UserId))
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/requests/" + rId + "/" + id + "/reject"
	response.HandleRequest(w, r, url)
}

func FinishAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	aId := params["appointmentId"]

	worker, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	id := strconv.Itoa(int(worker.UserId))
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/" + aId + "/" + id + "/finish"
	response.HandleRequest(w, r, url)
}

func CancelAppointment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	aId := params["appointmentId"]

	worker, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	id := strconv.Itoa(int(worker.UserId))
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/" + aId + "/" + id + "/cancel"
	response.HandleRequest(w, r, url)
}

func GetRequestsForCarService(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cId := params["carServiceId"]

	_, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/requests/service/" + cId
	response.HandleRequest(w, r, url)
}

func GetAppointmentsForCarService(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cId := params["carServiceId"]

	_, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/car-service/" + cId
	response.HandleRequest(w, r, url)
}

func GetAppointmentsForWorker(w http.ResponseWriter, r *http.Request) {
	worker, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	id := strconv.Itoa(int(worker.UserId))
	url := response.AppointmentServiceRoundRobin.Next().Host + "/api/appointments/worker/" + id
	response.HandleRequest(w, r, url)
}
