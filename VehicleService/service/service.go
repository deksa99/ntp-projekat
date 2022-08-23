package service

import (
	"VehicleService/model"
	"VehicleService/repository"
	"VehicleService/response"
	"VehicleService/util/converter"
	"errors"
)

func FindVehiclesForUser(userId uint) ([]response.VehicleInfo, error) {
	vehicles := repository.FindVehicles(userId)

	var vehicleInfos []response.VehicleInfo

	for _, v := range vehicles {
		vehicleInfos = append(vehicleInfos, converter.VehicleToVehicleInfo(&v))
	}

	return vehicleInfos, nil
}

func FindVehicleById(id uint) (response.VehicleInfo, error) {
	vehicle, err := repository.GetVehicleById(id)

	if err != nil {
		return response.VehicleInfo{}, err
	}

	vehicleInfo := converter.VehicleToVehicleInfo(&vehicle)

	return vehicleInfo, nil
}

func AddVehicle(userId uint, manufacturer model.Manufacturer, carModel string, color string, licencePlate string, chassisNumber string) (model.Vehicle, error) {

	vehicle := model.Vehicle{Color: color, CarModel: carModel, ChassisNumber: chassisNumber, LicencePlate: licencePlate, Manufacturer: manufacturer, UserID: userId}

	newVehicle, err := repository.Save(vehicle)

	if err != nil {
		return vehicle, err
	}

	return newVehicle, nil
}

func UpdateVehicle(userId uint, vehicleId uint, color string) (model.Vehicle, error) {
	vehicle, err := repository.GetVehicleById(vehicleId)
	if err != nil {
		return model.Vehicle{}, err
	}
	if vehicle.UserID != userId {
		return model.Vehicle{}, errors.New("oops... not your vehicle")
	}

	vehicle.Color = color

	_, err = repository.Save(vehicle)
	if err != nil {
		return model.Vehicle{}, err
	} else {
		return vehicle, nil
	}
}
