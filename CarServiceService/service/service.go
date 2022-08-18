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
