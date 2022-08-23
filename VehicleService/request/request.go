package request

import "VehicleService/model"

type AddVehicle struct {
	UserId        uint               `json:"UserId"`
	Manufacturer  model.Manufacturer `json:"Manufacturer"`
	CarModel      string             `json:"CarModel"`
	Color         string             `json:"Color"`
	LicencePlate  string             `json:"LicencePlate"`
	ChassisNumber string             `json:"ChassisNumber"`
}

type UpdateVehicle struct {
	Color string `json:"Color"`
}
