package auth

import (
	"UserService/response"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Validate(r *http.Request, roles []string) (response.Authentication, error) {
	header := r.Header.Values("Authorization")
	if len(header) == 0 {
		return response.Authentication{}, errors.New("header authorization not found")
	}

	tokenString := strings.Split(header[0], " ")
	if len(tokenString) != 2 {
		return response.Authentication{}, errors.New("invalid header format")
	}

	token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil || !token.Valid {
		return response.Authentication{}, err
	}

	claims := token.Claims.(jwt.MapClaims)

	exp, ok := claims["exp"].(float64)
	if !ok {
		return response.Authentication{}, errors.New("exp parsing error")
	}
	if time.Now().Unix() > int64(exp) {
		return response.Authentication{}, errors.New("token expired")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return response.Authentication{}, errors.New("role parsing error")
	}

	accId, ok := claims["id"].(float64)
	if !ok {
		return response.Authentication{}, errors.New("acc parsing error")
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return response.Authentication{}, errors.New("user parsing error")
	}

	if !roleInRoles(role, roles) {
		return response.Authentication{}, errors.New("forbidden")
	}

	return response.Authentication{AccId: uint(accId), UserId: uint(userId), Role: role}, nil
}

func roleInRoles(role string, roles []string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
