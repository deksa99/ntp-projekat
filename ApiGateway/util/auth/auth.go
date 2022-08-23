package auth

import (
	"ApiGateway/model"
	"ApiGateway/util/helper"
	"ApiGateway/util/response"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func Authentication(_ http.ResponseWriter, r *http.Request, roles []string) (model.AuthenticatedUser, error) {

	var authUser model.AuthenticatedUser

	aRoles := model.Roles{Roles: roles}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(aRoles)
	if err != nil {
		return authUser, errors.New("unauthorized")
	}

	authRequest, _ := http.NewRequest(http.MethodPost, response.UserServiceRoundRobin.Next().Host+"/api/users/auth", &buf)
	authRequest.Header.Set("Accept", "application/json")
	values := r.Header.Values("Authorization")

	if len(values) == 0 {
		return authUser, errors.New("unauthorized")
	}

	authRequest.Header.Set("Authorization", values[0])
	authClient := &http.Client{}
	authResponse, err := authClient.Do(authRequest)

	if err != nil {
		return authUser, errors.New("user service is not available")
	}

	if authResponse.StatusCode != 200 {
		return authUser, errors.New("unauthorized")
	}

	err = helper.GetJson(authResponse, &authUser)
	if err != nil {
		return authUser, errors.New("cannot convert to json")
	}

	return authUser, nil
}
