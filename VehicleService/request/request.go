package request

import "VehicleService/model"

type AddVehicle struct {
	Manufacturer  model.Manufacturer `json:"Manufacturer"`
	CarModel      string             `json:"CarModel"`
	Color         string             `json:"Color"`
	LicensePlate  string             `json:"LicensePlate"`
	ChassisNumber string             `json:"ChassisNumber"`
}

type UpdateVehicle struct {
	Color string `json:"Color"`
}
