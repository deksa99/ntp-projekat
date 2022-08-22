package handler

import (
	"ApiGateway/util/auth"
	"ApiGateway/util/response"
	"github.com/gorilla/mux"
	"net/http"
)

func ManufacturerReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["carServiceId"]

	_, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.ReportServiceRoundRobin.Next().Host + "/api/reports/manufacturer/" + id
	response.HandleRequest(w, r, url)
}

func ServiceReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["carServiceId"]

	_, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.ReportServiceRoundRobin.Next().Host + "/api/reports/service/" + id
	response.HandleRequest(w, r, url)
}

func FinancialReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["carServiceId"]

	_, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.ReportServiceRoundRobin.Next().Host + "/api/reports/financial/" + id
	response.HandleRequest(w, r, url)
}

func StatusReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["carServiceId"]

	_, err := auth.Authentication(w, r, []string{"worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.ReportServiceRoundRobin.Next().Host + "/api/reports/status/" + id
	response.HandleRequest(w, r, url)
}
