package service

import (
	"UserService/model"
	"UserService/repository"
	"time"
)

func CreateUser(Username string, Password string, FirstName string, LastName string, Email string) (model.User, error) {
	acc := model.Account{Username: Username, Password: Password, Active: true, BlockedUntil: time.Time{}}
	user := model.User{FirstName: FirstName, LastName: LastName, Email: Email, Account: acc}

	newUser, err := repository.CreateUser(user)

	if err != nil {
		return user, err
	}

	return newUser, nil
}
