package repository

import (
	"AppointmentService/model"
	"AppointmentService/util/database"
)

func SaveRequest(request *model.AppointmentRequest) (model.AppointmentRequest, error) {
	appRequest := database.Db.Save(&request)

	if appRequest.Error != nil {
		return *request, appRequest.Error
	}

	return *request, nil
}
