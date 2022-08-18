package service

import (
	"CarServiceService/model"
	"CarServiceService/repository"
	"CarServiceService/response"
	"CarServiceService/util/converter"
	"CarServiceService/util/helper"
)

func FindAllCarServices() []response.CarServiceInfo {
	services := repository.FindCarServices()

	var serviceInfos []response.CarServiceInfo

	for _, s := range services {
		serviceInfos = append(serviceInfos, converter.CarServiceToCarServiceInfo(&s))
	}

	return serviceInfos
}

func GetServicesForCarService(carServiceId uint) ([]response.ServiceInfo, error) {
	services, err := repository.FindServicesForCarService(carServiceId)

	if err != nil {
		return []response.ServiceInfo{}, err
	}

	var serviceInfos []response.ServiceInfo

	for _, s := range services {
		serviceInfos = append(serviceInfos, converter.ServiceToServiceInfo(&s))
	}

	return serviceInfos, nil
}

func FindNearestService(lat float64, lon float64) (model.CarService, error) {
	locations := repository.FindLocations()

	n, err := helper.Nearest(lat, lon, locations)

	if err != nil {
		return model.CarService{}, err
	}

	service, err := repository.FindServiceForLocation(n.ID)

	if err != nil {
		return model.CarService{}, err
	} else {
		return service, nil
	}
}

func CreateService(carServiceId uint, name string, description string, price float32, expectedTIme uint) (response.ServiceInfo, error) {
	service := model.Service{
		Name:         name,
		Description:  description,
		Price:        price,
		ExpectedTime: expectedTIme,
		Available:    true,
		CarServiceID: carServiceId,
	}

	newService, err := repository.Save(service)

	if err != nil {
		return response.ServiceInfo{}, err
	} else {
		return converter.ServiceToServiceInfo(&newService), nil
	}
}
