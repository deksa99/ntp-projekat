package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Validate(r *http.Request, roles []string) (uint, error) {
	header := r.Header.Values("Authorization")
	if len(header) == 0 {
		return 0, errors.New("header authorization not found")
	}

	tokenString := strings.Split(header[0], " ")
	if len(tokenString) != 2 {
		return 0, errors.New("invalid header format")
	}

	token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)

	exp, ok := claims["exp"].(float64)
	if !ok {
		return 0, errors.New("exp parsing error")
	}
	if time.Now().Unix() > int64(exp) {
		return 0, errors.New("token expired")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return 0, errors.New("role parsing error")
	}

	accId, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("acc parsing error")
	}

	if !roleInRoles(role, roles) {
		return 0, errors.New("forbidden")
	}

	return uint(accId), nil
}

func roleInRoles(role string, roles []string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
