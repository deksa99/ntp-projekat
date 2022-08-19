package service

import (
	"UserService/model"
	"UserService/repository"
	"UserService/response"
	"UserService/util/converter"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username string, password string, firstName string, lastName string, email string) (model.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return model.User{}, err
	}

	acc := model.Account{Username: username, Password: string(hashedPassword), Active: true, BlockedUntil: time.Time{}}
	user := model.User{FirstName: firstName, LastName: lastName, Email: email, Account: acc}

	newUser, err := repository.CreateUser(user)

	if err != nil {
		return user, err
	}

	return newUser, nil
}

func Login(username string, password string, role string) (response.Jwt, error) {
	acc, err := repository.FindAccountByUsername(username)

	if err != nil {
		return response.Jwt{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))

	if err != nil {
		return response.Jwt{}, errors.New("invalid username or password")
	}

	if !acc.Active {
		return response.Jwt{}, errors.New("account is not active")
	}

	if time.Now().Before(acc.BlockedUntil) {
		return response.Jwt{}, errors.New("account blocked until: " + acc.BlockedUntil.String())
	}

	switch role {
	case "user":
		{
			u, err := repository.FindUserByAccountID(acc.ID)
			if err != nil {
				return response.Jwt{}, err
			} else {
				return createJwt(acc, "user", u.ID), nil
			}
		}
	case "admin":
		{
			a, err := repository.FindAdminByAccountID(acc.ID)
			if err != nil {
				return response.Jwt{}, err
			} else {
				return createJwt(acc, "admin", a.ID), nil
			}
		}
	case "worker":
		{
			w, err := repository.FindWorkerByAccountID(acc.ID)
			if err != nil {
				return response.Jwt{}, err
			} else {
				return createJwt(acc, "worker", w.ID), nil
			}
		}
	default:
		return response.Jwt{}, errors.New("unsupported role")
	}
}

func createJwt(acc model.Account, role string, userId uint) response.Jwt {

	jwtKey := []byte(os.Getenv("JWT_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := jwt.MapClaims{}
	claims["username"] = acc.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["role"] = role
	claims["id"] = acc.ID
	claims["user_id"] = userId

	token.Claims = claims

	tokenString, _ := token.SignedString(jwtKey)

	jwtR := response.Jwt{Token: tokenString}

	return jwtR
}

func ChangePassword(accId uint, password string, newPassword string) error {
	acc, err := repository.FindAccountById(accId)

	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))

	if err != nil {
		return errors.New("invalid username or password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		return err
	}

	err = repository.UpdatePassword(accId, string(hashedPassword))

	if err != nil {
		return err
	}

	return nil
}

func FindAllUsers() []response.UserInfo {
	users := repository.FindUsers()

	var userInfos []response.UserInfo

	for _, u := range users {
		userInfos = append(userInfos, converter.UserToUserInfo(&u))
	}

	return userInfos
}

func FindUser(id uint) (response.UserInfo, error) {
	user, err := repository.FindUserById(id)

	if err != nil {
		return response.UserInfo{}, err
	}

	ui := converter.UserToUserInfo(&user)

	return ui, nil
}

func FindWorker(id uint) (response.WorkerInfo, error) {
	worker, err := repository.FindWorkerById(id)

	if err != nil {
		return response.WorkerInfo{}, err
	}

	wi := converter.WorkerToWorkerInfo(&worker)

	return wi, nil
}

func BlockUser(accId uint) error {
	err := repository.BlockUser(accId)

	if err != nil {
		return err
	}

	return nil
}
