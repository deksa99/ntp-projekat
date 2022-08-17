package service

import (
	"VehicleService/repository"
	"VehicleService/response"
	"VehicleService/util/converter"
)

func FindVehiclesForUser(userId uint) ([]response.VehicleInfo, error) {
	vehicles := repository.FindVehicles(userId)

	var vehicleInfos []response.VehicleInfo

	for _, v := range vehicles {
		vehicleInfos = append(vehicleInfos, converter.VehicleToVehicleInfo(&v))
	}

	return vehicleInfos, nil
}
