package repository

import (
	"UserService/model"
	"UserService/util/database"
	"errors"
	"net/mail"
)

func CreateUser(user model.User) (model.User, error) {
	_, err := mail.ParseAddress(user.Email)

	if err != nil {
		return user, errors.New("invalid email format")
	}

	newUser := database.Db.Create(&user)

	if newUser.Error != nil {
		return user, newUser.Error
	}

	return user, nil
}
