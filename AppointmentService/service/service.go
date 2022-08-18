package service

import (
	"AppointmentService/model"
	"AppointmentService/repository"
	"AppointmentService/response"
	"AppointmentService/util/api"
	"AppointmentService/util/converter"
	"errors"
	"time"
)

func CreateAppointmentRequest(userId uint, vehicleId uint, serviceId uint, carServiceId uint) (response.AppointmentRequestInfo, error) {
	user, err := api.GetUserDetails(userId)
	if err != nil {
		return response.AppointmentRequestInfo{}, err
	}
	vehicle, err := api.GetVehicleDetails(vehicleId)
	if err != nil {
		return response.AppointmentRequestInfo{}, err
	}
	service, err := api.GetService(serviceId)
	if err != nil {
		return response.AppointmentRequestInfo{}, err
	}
	carService, err := api.GetCarService(carServiceId)
	if err != nil {
		return response.AppointmentRequestInfo{}, err
	}

	if carServiceId != service.CarServiceID {
		return response.AppointmentRequestInfo{}, errors.New("service does not belong to the given car service")
	}

	if userId != vehicle.UserID {
		return response.AppointmentRequestInfo{}, errors.New("vehicle does not belong to the given user")
	}

	request := model.AppointmentRequest{
		UserID:         userId,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		VehicleID:      vehicleId,
		Manufacturer:   vehicle.Manufacturer,
		CarModel:       vehicle.CarModel,
		LicencePlate:   vehicle.LicencePlate,
		ChassisNumber:  vehicle.ChassisNumber,
		ServiceID:      serviceId,
		ServiceName:    service.ServiceName,
		ServicePrice:   service.ServicePrice,
		Available:      service.Available,
		CarServiceID:   carServiceId,
		CarServiceName: carService.CarServiceName,
		Street:         carService.Street,
		Status:         model.Submitted,
		SubmittedAt:    time.Now(),
	}

	newRequest, err := repository.SaveRequest(&request)

	if err != nil {
		return response.AppointmentRequestInfo{}, err
	}

	return converter.AppointmentRequestToAppointmentRequestInfo(&newRequest), nil
}
