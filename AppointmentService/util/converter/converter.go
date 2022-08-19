package converter

import (
	"AppointmentService/model"
	"AppointmentService/response"
)

func AppointmentRequestToAppointmentRequestInfo(ar *model.AppointmentRequest) response.AppointmentRequestInfo {
	return response.AppointmentRequestInfo{
		UserId:         ar.UserID,
		FirstName:      ar.FirstName,
		LastName:       ar.LastName,
		Email:          ar.Email,
		VehicleId:      ar.VehicleID,
		Manufacturer:   ar.Manufacturer,
		CarModel:       ar.CarModel,
		LicencePlate:   ar.LicencePlate,
		ChassisNumber:  ar.ChassisNumber,
		ServiceId:      ar.ServiceID,
		ServiceName:    ar.ServiceName,
		ServicePrice:   ar.ServicePrice,
		Available:      ar.Available,
		CarServiceId:   ar.CarServiceID,
		CarServiceName: ar.CarServiceName,
		Street:         ar.Street,
		Status:         ar.Status,
		SubmittedAt:    ar.SubmittedAt,
	}
}

func AppointmentToAppointmentInfo(a *model.Appointment) response.AppointmentInfo {
	return response.AppointmentInfo{
		UserId:         a.AppointmentRequest.UserID,
		FirstName:      a.AppointmentRequest.FirstName,
		LastName:       a.AppointmentRequest.LastName,
		Email:          a.AppointmentRequest.Email,
		VehicleId:      a.AppointmentRequest.VehicleID,
		Manufacturer:   a.AppointmentRequest.Manufacturer,
		CarModel:       a.AppointmentRequest.CarModel,
		LicencePlate:   a.AppointmentRequest.LicencePlate,
		ChassisNumber:  a.AppointmentRequest.ChassisNumber,
		ServiceId:      a.AppointmentRequest.ServiceID,
		ServiceName:    a.AppointmentRequest.ServiceName,
		ServicePrice:   a.AppointmentRequest.ServicePrice,
		CarServiceId:   a.AppointmentRequest.CarServiceID,
		CarServiceName: a.AppointmentRequest.CarServiceName,
		Street:         a.AppointmentRequest.Street,
		StartTime:      a.StartTime,
		FinishTime:     a.FinishTime,
		Status:         a.Status,
	}
}
