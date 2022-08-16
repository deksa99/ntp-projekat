package service

import (
	"UserService/model"
	"UserService/repository"
	"UserService/response"
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
			_, err := repository.FindUserByAccountID(acc.ID)
			if err != nil {
				return response.Jwt{}, err
			} else {
				return createJwt(acc, "user"), nil
			}
		}
	case "admin":
		{
			_, err := repository.FindAdminByAccountID(acc.ID)
			if err != nil {
				return response.Jwt{}, err
			} else {
				return createJwt(acc, "admin"), nil
			}
		}
	case "worker":
		{
			_, err := repository.FindWorkerByAccountID(acc.ID)
			if err != nil {
				return response.Jwt{}, err
			} else {
				return createJwt(acc, "worker"), nil
			}
		}
	default:
		return response.Jwt{}, errors.New("unsupported role")
	}
}

func createJwt(acc model.Account, role string) response.Jwt {

	jwtKey := []byte(os.Getenv("JWT_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := jwt.MapClaims{}
	claims["username"] = acc.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["role"] = role
	claims["id"] = acc.ID

	token.Claims = claims

	tokenString, _ := token.SignedString(jwtKey)

	jwt := response.Jwt{Token: tokenString}

	return jwt
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
