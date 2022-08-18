package service

import (
	"CarServiceService/repository"
	"CarServiceService/response"
	"CarServiceService/util/converter"
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
