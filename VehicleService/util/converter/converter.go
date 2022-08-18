package converter

import (
	"VehicleService/model"
	"VehicleService/response"
)

func VehicleToVehicleInfo(vehicle *model.Vehicle) response.VehicleInfo {
	return response.VehicleInfo{
		Id:            vehicle.ID,
		Manufacturer:  vehicle.Manufacturer,
		CarModel:      vehicle.CarModel,
		Color:         vehicle.Color,
		LicensePlate:  vehicle.LicencePlate,
		ChassisNumber: vehicle.ChassisNumber,
		UserId:        vehicle.UserID,
	}
}
