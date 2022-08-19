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

func GetRequestsForUser(id uint) []model.AppointmentRequest {
	var appRequests []model.AppointmentRequest
	database.Db.Where("user_id = ?", id).Find(&appRequests)

	return appRequests
}

func GetAppointmentsForUser(id uint) []model.Appointment {
	var appointments []model.Appointment
	database.Db.Joins("AppointmentRequest").Where("user_id = ?", id).Find(&appointments)
	return appointments
}
