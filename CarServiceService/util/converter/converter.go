package converter

import (
	"CarServiceService/model"
	"CarServiceService/response"
)

func CarServiceToCarServiceInfo(carService *model.CarService) response.CarServiceInfo {
	return response.CarServiceInfo{
		Id:        carService.ID,
		Name:      carService.Name,
		Longitude: carService.Location.Longitude,
		Latitude:  carService.Location.Latitude,
		Street:    carService.Location.Street,
	}
}
