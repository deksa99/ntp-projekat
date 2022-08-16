package handler

import (
	"UserService/request"
	"UserService/response"
	"UserService/service"
	"UserService/util/converter"
	"UserService/util/helper"
	"errors"
	"log"

	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest request.Login

	err := helper.DecodeJSONBody(w, r, &loginRequest)
	role := r.URL.Query().Get("role")

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		var e *response.Error
		if errors.As(err, &e) {
			http.Error(w, e.Message, e.Status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	jwt, err := service.Login(loginRequest.Username, loginRequest.Password, role)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		err = json.NewEncoder(w).Encode(jwt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

}

func Register(w http.ResponseWriter, r *http.Request) {
	var userRequest request.CreateUser

	err := helper.DecodeJSONBody(w, r, &userRequest)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		var e *response.Error
		if errors.As(err, &e) {
			http.Error(w, e.Message, e.Status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	user, err := service.CreateUser(
		userRequest.Username,
		userRequest.Password,
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Email,
	)

	if err != nil {
		err = json.NewEncoder(w).Encode(response.Error{Message: err.Error(), Status: http.StatusBadRequest})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(converter.UserToUserInfo(&user))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {

}

func FindUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {

}

func BlockUser(w http.ResponseWriter, r *http.Request) {

}
