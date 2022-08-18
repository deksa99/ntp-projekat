package response

import "VehicleService/model"

type Error struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

func (e *Error) Error() string {
	return e.Message
}

type VehicleInfo struct {
	Id            uint               `json:"Id"`
	UserId        uint               `json:"UserId"`
	Manufacturer  model.Manufacturer `json:"Manufacturer"`
	CarModel      string             `json:"CarModel"`
	Color         string             `json:"Color"`
	LicencePlate  string             `json:"LicencePlate"`
	ChassisNumber string             `json:"ChassisNumber"`
}
