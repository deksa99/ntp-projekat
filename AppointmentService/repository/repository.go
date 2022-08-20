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

func SaveAppointment(appointment model.Appointment) (model.Appointment, error) {
	newAppointment := database.Db.Save(&appointment)

	if newAppointment.Error != nil {
		return appointment, newAppointment.Error
	}

	database.Db.Preload("AppointmentRequest").First(&appointment, appointment.ID)

	return appointment, nil
}

func FindRequestById(requestId uint) (model.AppointmentRequest, error) {
	var appRequest model.AppointmentRequest
	database.Db.First(&appRequest, requestId)

	if appRequest.ID == 0 {
		return appRequest, errors.New("appointment request not found")
	}

	return appRequest, nil
}

func FindAppointmentById(appointmentId uint) (model.Appointment, error) {
	var app model.Appointment
	database.Db.Preload("AppointmentRequest").First(&app, appointmentId)

	if app.ID == 0 {
		return app, errors.New("appointment not found")
	}

	return app, nil
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

func GetAppointmentsForWorker(id uint) []model.Appointment {
	var appointments []model.Appointment
	database.Db.Joins("AppointmentRequest").Where("worker_id = ?", id).Find(&appointments)
	return appointments
}

func GetRequestsForCarService(id uint) []model.AppointmentRequest {
	var appRequests []model.AppointmentRequest
	database.Db.Where("car_service_id = ?", id).Find(&appRequests)

	return appRequests
}
