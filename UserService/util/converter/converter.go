package converter

import (
	"UserService/model"
	"UserService/response"
)

func UserToUserInfo(user *model.User) response.UserInfo {
	return response.UserInfo{
		Id:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName}
}
