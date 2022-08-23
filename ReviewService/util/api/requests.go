package api

import (
	"encoding/json"
	"errors"
	rr "github.com/hlts2/round-robin"
	"net/http"
	"net/url"
	"strconv"
)

var appointmentServiceBasePath, _ = rr.New(
	&url.URL{Host: "http://localhost:8084"},
)

func GetAppointmentDetails(id uint) (AppointmentInfo, error) {
	response, err := http.Get(appointmentServiceBasePath.Next().Host + "/api/appointments/" + strconv.FormatUint(uint64(id), 10))

	if err != nil {
		return AppointmentInfo{}, err
	}
	if response.StatusCode != 200 {
		return AppointmentInfo{}, errors.New("appointment not found")
	}

	var app AppointmentInfo
	err = json.NewDecoder(response.Body).Decode(&app)
	if err != nil {
		return app, err
	} else {
		return app, nil
	}
}
