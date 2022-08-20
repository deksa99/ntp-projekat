package handler

import (
	"ReviewService/request"
	"ReviewService/response"
	"ReviewService/service"
	"ReviewService/util/helper"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetReportedReviews(w http.ResponseWriter, _ *http.Request) {
	reported, err := service.GetReportedReviews()

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		err = json.NewEncoder(w).Encode(reported)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	var cr request.AddReview

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

	newReview, err := service.AddReview(cr.AppointmentID, cr.ServiceID, cr.CarServiceID, cr.UserID, cr.Rating, cr.Comment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		err = json.NewEncoder(w).Encode(newReview)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func ReportReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["reviewId"], 10, 32)

	reported, err := service.ReportReview(uint(id))

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		err = json.NewEncoder(w).Encode(reported)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func ProcessReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["reportId"], 10, 32)

	var pr request.ProcessReport

	err := helper.DecodeJSONBody(w, r, &pr)

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

	reported, err := service.ProcessReport(uint(id), pr.Inappropriate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		err = json.NewEncoder(w).Encode(reported)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
