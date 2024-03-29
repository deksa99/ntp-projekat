package service

import (
	"CarServiceService/model"
	"CarServiceService/repository"
	"CarServiceService/response"
	"CarServiceService/util/converter"
	"CarServiceService/util/helper"
)

func GetCarService(id uint) (response.CarServiceInfo, error) {
	carService, err := repository.FindCarServiceById(id)

	if err != nil {
		return response.CarServiceInfo{}, err
	}

	carServiceInfo := converter.CarServiceToCarServiceInfo(&carService)
	return carServiceInfo, nil
}

func GetService(id uint) (response.ServiceInfo, error) {
	service, err := repository.FindServiceById(id)

	if err != nil {
		return response.ServiceInfo{}, err
	}

	serviceInfo := converter.ServiceToServiceInfo(&service)
	return serviceInfo, nil
}

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

	service, err := repository.FindCarServiceForLocation(n.ID)

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

	newService, err := repository.SaveService(service)

	if err != nil {
		return response.ServiceInfo{}, err
	} else {
		return converter.ServiceToServiceInfo(&newService), nil
	}
}

func UpdateService(id uint, name string, description string, expectedTIme uint) (response.ServiceInfo, error) {
	service, err := repository.FindServiceById(id)

	service.Name = name
	service.Description = description
	service.ExpectedTime = expectedTIme

	newService, err := repository.SaveService(service)

	if err != nil {
		return response.ServiceInfo{}, err
	} else {
		return converter.ServiceToServiceInfo(&newService), nil
	}
}

func ChangeServiceAvailability(id uint) (response.ServiceInfo, error) {
	service, err := repository.FindServiceById(id)

	service.Available = !service.Available

	newService, err := repository.SaveService(service)

	if err != nil {
		return response.ServiceInfo{}, err
	} else {
		return converter.ServiceToServiceInfo(&newService), nil
	}
}

func CreateCarService(name string, latitude float64, longitude float64, street string) (response.CarServiceInfo, error) {
	location := model.Location{
		Latitude:  latitude,
		Longitude: longitude,
		Street:    street,
	}

	carService := model.CarService{
		Name:     name,
		Location: location,
	}

	newCarService, err := repository.SaveCarService(carService)

	if err != nil {
		return response.CarServiceInfo{}, err
	} else {
		return converter.CarServiceToCarServiceInfo(&newCarService), nil
	}
}
