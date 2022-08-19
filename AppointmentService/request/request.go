package request

import "time"

type CreateAppointmentRequest struct {
	UserID       uint `json:"UserID"`
	VehicleID    uint `json:"VehicleID"`
	ServiceID    uint `json:"ServiceID"`
	CarServiceID uint `json:"CarServiceID"`
}

type CancelAppointmentRequest struct {
	UserID               uint `json:"UserID"`
	AppointmentRequestID uint `json:"AppointmentRequestID"`
}

type AcceptAppointmentRequest struct {
	Start time.Time `json:"Start"`
}
