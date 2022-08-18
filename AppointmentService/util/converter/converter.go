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
