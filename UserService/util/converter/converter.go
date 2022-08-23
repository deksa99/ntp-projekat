package converter

import (
	"UserService/model"
	"UserService/response"
)

func UserToUserInfo(user *model.User) response.UserInfo {
	return response.UserInfo{
		Id:        user.ID,
		Email:     user.Email,
		Username:  user.Account.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName}
}

func WorkerToWorkerInfo(worker *model.ServiceWorker) response.WorkerInfo {
	return response.WorkerInfo{
		WorkerID:     worker.ID,
		FirstName:    worker.FirstName,
		LastName:     worker.LastName,
		CarServiceID: worker.CarServiceID,
	}
}
