package repository

import (
	"VehicleService/model"
	"VehicleService/util/database"
)

func FindVehicles(id uint) []model.Vehicle {
	var vehicles []model.Vehicle

	database.Db.Table("vehicles").Where("user_id = ?", id).Find(&vehicles)

	return vehicles
}
