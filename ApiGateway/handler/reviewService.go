package handler

import (
	"ApiGateway/util/auth"
	"ApiGateway/util/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddReview(w http.ResponseWriter, r *http.Request) {
	u, err := auth.Authentication(w, r, []string{"user"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.ReviewServiceRoundRobin.Next().Host + "/api/reviews/add/" + strconv.Itoa(int(u.UserId))
	response.HandleRequest(w, r, url)
}

func ReportReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := params["id"]

	_, err := auth.Authentication(w, r, []string{"user", "worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.ReviewServiceRoundRobin.Next().Host + "/api/reviews/" + id + "/report"
	response.HandleRequest(w, r, url)
}

func ProcessReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := params["id"]

	_, err := auth.Authentication(w, r, []string{"admin"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.ReviewServiceRoundRobin.Next().Host + "/api/reviews/report/" + id + "/process"
	response.HandleRequest(w, r, url)
}

func GetReports(w http.ResponseWriter, r *http.Request) {
	_, err := auth.Authentication(w, r, []string{"admin"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.ReviewServiceRoundRobin.Next().Host + "/api/reviews/reported"
	response.HandleRequest(w, r, url)
}

func GetReviews(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := params["carServiceId"]

	url := response.ReviewServiceRoundRobin.Next().Host + "/api/reviews/car-service/" + id
	response.HandleRequest(w, r, url)
}
