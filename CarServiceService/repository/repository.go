package repository

import (
	"CarServiceService/model"
	"CarServiceService/util/database"
)

func FindCarServices() []model.CarService {
	var carServices []model.CarService

	database.Db.Preload("Location").Find(&carServices)

	return carServices
}
