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

	newRequest, err := repository.SaveRequest(request)

	if err != nil {
		return response.AppointmentRequestInfo{}, err
	}

	return converter.AppointmentRequestToAppointmentRequestInfo(&newRequest), nil
}

func CancelAppointmentRequest(userId uint, appointmentRequestId uint) (response.AppointmentRequestInfo, error) {
	request, err := repository.FindRequestById(appointmentRequestId)
	if err != nil {
		return response.AppointmentRequestInfo{}, err
	}

	if userId != request.UserID {
		return response.AppointmentRequestInfo{}, errors.New("oops you can only cancel your request")
	}

	if request.Status != model.Submitted {
		return response.AppointmentRequestInfo{}, errors.New("request has already been " + string(request.Status))
	}

	if !request.Available {
		return response.AppointmentRequestInfo{}, errors.New("service '" + request.ServiceName + "' not available")
	}

	request.Status = model.CancelledRequest

	cancelledRequest, err := repository.SaveRequest(request)
	if err != nil {
		return response.AppointmentRequestInfo{}, err
	}
	return converter.AppointmentRequestToAppointmentRequestInfo(&cancelledRequest), nil
}

func GetRequestsForUser(id uint) ([]response.AppointmentRequestInfo, error) {
	requests := repository.GetRequestsForUser(id)

	var requestInfos []response.AppointmentRequestInfo

	for _, r := range requests {
		requestInfos = append(requestInfos, converter.AppointmentRequestToAppointmentRequestInfo(&r))
	}

	if requestInfos == nil {
		return []response.AppointmentRequestInfo{}, nil
	}

	return requestInfos, nil
}

func GetAppointmentsForUser(id uint) ([]response.AppointmentInfo, error) {
	appointments := repository.GetAppointmentsForUser(id)

	var requestInfos []response.AppointmentInfo

	for _, a := range appointments {
		requestInfos = append(requestInfos, converter.AppointmentToAppointmentInfo(&a))
	}

	if requestInfos == nil {
		return []response.AppointmentInfo{}, nil
	}

	return requestInfos, nil
}

func AcceptRequest(requestId uint, workerID uint, start time.Time) (response.AppointmentInfo, error) {
	worker, err := api.GetWorkerDetails(workerID)
	if err != nil {
		return response.AppointmentInfo{}, err
	}
	request, err := repository.FindRequestById(requestId)
	if err != nil {
		return response.AppointmentInfo{}, err
	}

	if request.Status != model.Submitted {
		return response.AppointmentInfo{}, errors.New("request has already been " + string(request.Status))
	}

	if worker.CarServiceID != request.CarServiceID {
		return response.AppointmentInfo{}, errors.New("worker does not work here")
	}

	appointment := model.Appointment{
		AppointmentRequestID: request.ID,
		WorkerID:             worker.WorkerID,
		FirstName:            worker.FirstName,
		LastName:             worker.LastName,
		StartTime:            start,
		FinishTime:           time.Time{},
		Status:               model.Scheduled,
	}

	request.Status = model.Accepted
	_, err = repository.SaveRequest(request)
	if err != nil {
		return response.AppointmentInfo{}, err
	}

	newAppointment, err := repository.SaveAppointment(appointment)
	newAppointment.AppointmentRequest = request
	if err != nil {
		return response.AppointmentInfo{}, err
	} else {
		return converter.AppointmentToAppointmentInfo(&newAppointment), nil
	}
}

func RejectRequest(requestId uint, workerID uint) (response.AppointmentInfo, error) {
	return response.AppointmentInfo{}, errors.New("not implemented")
}
