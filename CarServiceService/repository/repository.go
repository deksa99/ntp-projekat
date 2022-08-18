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

func FindLocations() []model.Location {
	var locations []model.Location

	database.Db.Table("locations").Find(&locations)

	return locations
}

func FindServiceForLocation(lId uint) (model.CarService, error) {
	var carService model.CarService

	database.Db.Preload("Location").Table("car_services").Where("location_id = ?", lId).First(&carService)

	if carService.ID == 0 {
		return carService, errors.New("car service not found")
	}

	return carService, nil
}

func Save(service model.Service) (model.Service, error) {
	newVehicle := database.Db.Save(&service)

	if newVehicle.Error != nil {
		return service, newVehicle.Error
	}

	return service, nil
}
