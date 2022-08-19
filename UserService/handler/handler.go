package handler

import (
	"UserService/request"
	"UserService/response"
	"UserService/service"
	"UserService/util/auth"
	"UserService/util/converter"
	"UserService/util/helper"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"strconv"

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
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	var changePassword request.ChangePassword

	err := helper.DecodeJSONBody(w, r, &changePassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.ChangePassword(uint(id), changePassword.Password, changePassword.NewPassword)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	user, err := service.FindUser(uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func FindWorker(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	w.Header().Set("Content-Type", "application/json")

	user, err := service.FindWorker(uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func FindAllUsers(w http.ResponseWriter, _ *http.Request) {
	users := service.FindAllUsers()

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	roles := []string{"admin"}
	_, err := auth.Validate(r, roles)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	err = service.BlockUser(uint(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var authRequest request.Authenticate

	err := helper.DecodeJSONBody(w, r, &authRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authRes, err := auth.Validate(r, authRequest.Roles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	err = json.NewEncoder(w).Encode(authRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
