package response

import "VehicleService/model"

type VehicleInfo struct {
	Id            uint               `json:"Id"`
	Manufacturer  model.Manufacturer `json:"Manufacturer"`
	CarModel      string             `json:"CarModel"`
	Color         string             `json:"Color"`
	LicensePlate  string             `json:"LicensePlate"`
	ChassisNumber string             `json:"ChassisNumber"`
}
