package repository

import (
	"UserService/model"
	"UserService/util/database"
	"errors"
	"net/mail"
	"time"
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

func FindAccountByUsername(username string) (model.Account, error) {
	var acc model.Account

	database.Db.Table("accounts").Where("username = ?", username).First(&acc)

	if acc.ID == 0 {
		return acc, errors.New("user " + username + " not found")
	}

	return acc, nil
}

func FindAccountById(id uint) (model.Account, error) {
	var acc model.Account

	database.Db.Table("accounts").Where("id = ?", id).First(&acc)

	if acc.ID == 0 {
		return acc, errors.New("account not found")
	}

	return acc, nil
}

func FindUserByAccountID(accountId uint) (model.User, error) {
	var user model.User

	database.Db.Table("users").Where("account_id = ?", accountId).First(&user)

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

func FindUsers() []model.User {
	var users []model.User

	database.Db.Preload("Account").Find(&users)

	return users
}

func FindUserById(id uint) (model.User, error) {
	var user model.User

	database.Db.Preload("Account").First(&user, id)

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

func FindWorkerById(id uint) (model.ServiceWorker, error) {
	var worker model.ServiceWorker

	database.Db.First(&worker, id)

	if worker.ID == 0 {
		return worker, errors.New("worker not found")
	}

	return worker, nil
}

func FindAdminByAccountID(accountId uint) (model.Admin, error) {
	var admin model.Admin

	database.Db.Table("admins").Where("account_id = ?", accountId).First(&admin)

	if admin.ID == 0 {
		return admin, errors.New("admin not found")
	}

	return admin, nil
}

func FindWorkerByAccountID(accountId uint) (model.ServiceWorker, error) {
	var worker model.ServiceWorker

	database.Db.Table("service_workers").Where("account_id = ?", accountId).First(&worker)

	if worker.ID == 0 {
		return worker, errors.New("worker not found")
	}

	return worker, nil
}

func UpdatePassword(id uint, newPassword string) error {
	var acc model.Account

	database.Db.First(&acc, id)

	if acc.ID == 0 {
		return errors.New("account not found")
	}

	acc.Password = newPassword
	database.Db.Save(&acc)

	return nil
}

func BlockUser(id uint) error {
	var acc model.Account

	database.Db.First(&acc, id)

	if acc.ID == 0 {
		return errors.New("account not found")
	}

	acc.BlockedUntil = time.Now().Add(time.Hour * 72)
	database.Db.Save(&acc)

	return nil
}
