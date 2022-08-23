package repository

import (
	"VehicleService/model"
	"VehicleService/util/database"
	"errors"
)

func FindVehicles(id uint) []model.Vehicle {
	var vehicles []model.Vehicle

	database.Db.Table("vehicles").Where("user_id = ?", id).Find(&vehicles)

	return vehicles
}

func Save(vehicle model.Vehicle) (model.Vehicle, error) {
	newVehicle := database.Db.Save(&vehicle)

	if newVehicle.Error != nil {
		return vehicle, newVehicle.Error
	}

	return vehicle, nil
}

func GetVehicleById(id uint) (model.Vehicle, error) {
	var vehicle model.Vehicle

	database.Db.Table("vehicles").Where("id = ?", id).First(&vehicle)

	if vehicle.ID == 0 {
		return vehicle, errors.New("vehicle does not exist")
	}

	return vehicle, nil
}
