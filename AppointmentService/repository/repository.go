package repository

import (
	"AppointmentService/model"
	"AppointmentService/util/database"
	"errors"
)

func SaveRequest(request model.AppointmentRequest) (model.AppointmentRequest, error) {
	appRequest := database.Db.Save(&request)

	if appRequest.Error != nil {
		return request, appRequest.Error
	}

	return request, nil
}

func FindRequestById(requestId uint) (model.AppointmentRequest, error) {
	var appRequest model.AppointmentRequest
	database.Db.First(&appRequest, requestId)

	if appRequest.ID == 0 {
		return appRequest, errors.New("appointment request not found")
	}

	return appRequest, nil
}
