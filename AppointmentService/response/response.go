package response

import (
	"AppointmentService/model"
	"time"
)

type Error struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

func (e *Error) Error() string {
	return e.Message
}

type AppointmentRequestInfo struct {
	UserId         uint                `json:"UserId"`
	FirstName      string              `json:"FirstName"`
	LastName       string              `json:"LastName"`
	Email          string              `json:"Email"`
	VehicleId      uint                `json:"VehicleId"`
	Manufacturer   string              `json:"Manufacturer"`
	CarModel       string              `json:"CarModel"`
	LicencePlate   string              `json:"LicencePlate"`
	ChassisNumber  string              `json:"ChassisNumber"`
	ServiceId      uint                `json:"ServiceId"`
	ServiceName    string              `json:"ServiceName"`
	ServicePrice   float32             `json:"ServicePrice"`
	Available      bool                `json:"Available"`
	CarServiceId   uint                `json:"CarServiceId"`
	CarServiceName string              `json:"CarServiceName"`
	Street         string              `json:"Street"`
	Status         model.RequestStatus `json:"Status"`
	SubmittedAt    time.Time           `json:"SubmittedAt"`
}
