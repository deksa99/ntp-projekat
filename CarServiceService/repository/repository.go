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

func FindCarServiceForLocation(lId uint) (model.CarService, error) {
	var carService model.CarService

	database.Db.Preload("Location").Table("car_services").Where("location_id = ?", lId).First(&carService)

	if carService.ID == 0 {
		return carService, errors.New("car service not found")
	}

	return carService, nil
}

func SaveService(service model.Service) (model.Service, error) {
	newService := database.Db.Save(&service)

	if newService.Error != nil {
		return service, newService.Error
	}

	return service, nil
}

func SaveCarService(carService model.CarService) (model.CarService, error) {
	newCarService := database.Db.Save(&carService)

	if newCarService.Error != nil {
		return carService, newCarService.Error
	}

	return carService, nil
}

func FindServiceById(id uint) (model.Service, error) {
	var service model.Service

	database.Db.Table("services").Where("id = ?", id).First(&service)

	if service.ID == 0 {
		return service, errors.New("service not found")
	}

	return service, nil
}
