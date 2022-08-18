package repository

import (
	"CarServiceService/model"
	"CarServiceService/util/database"
	"errors"
)

func FindCarServices() []model.CarService {
	var carServices []model.CarService

	database.Db.Preload("Location").Find(&carServices)

	return carServices
}

func FindServicesForCarService(carServiceId uint) ([]model.Service, error) {
	var carService model.CarService

	database.Db.Preload("Services").Where("id = ?", carServiceId).First(&carService)

	if carService.ID == 0 {
		return []model.Service{}, errors.New("car service not found")
	} else {
		return carService.Services, nil
	}
}
