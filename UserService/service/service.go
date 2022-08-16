package service

import (
	"UserService/model"
	"UserService/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func CreateUser(Username string, Password string, FirstName string, LastName string, Email string) (model.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), 12)
	if err != nil {
		return model.User{}, err
	}

	acc := model.Account{Username: Username, Password: string(hashedPassword), Active: true, BlockedUntil: time.Time{}}
	user := model.User{FirstName: FirstName, LastName: LastName, Email: Email, Account: acc}

	newUser, err := repository.CreateUser(user)

	if err != nil {
		return user, err
	}

	return newUser, nil
}
