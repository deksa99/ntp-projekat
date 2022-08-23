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

func ServiceToServiceInfo(service *model.Service) response.ServiceInfo {
	return response.ServiceInfo{
		Id:           service.ID,
		Name:         service.Name,
		Price:        service.Price,
		Description:  service.Description,
		ExpectedTime: service.ExpectedTime,
		Available:    service.Available,
		CarServiceID: service.CarServiceID,
	}
}
