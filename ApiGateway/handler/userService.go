package handler

import (
	"ApiGateway/util/auth"
	"ApiGateway/util/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	role := r.URL.Query().Get("role")
	url := response.UserServiceRoundRobin.Next().Host + "/api/users/login?role=" + role

	response.HandleRequest(w, r, url)
}

func Register(w http.ResponseWriter, r *http.Request) {
	url := response.UserServiceRoundRobin.Next().Host + "/api/users/register"

	response.HandleRequest(w, r, url)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	resp, err := auth.Authentication(w, r, []string{"user", "admin", "worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	id := resp.AccId
	url := response.UserServiceRoundRobin.Next().Host + "/api/users/" + strconv.Itoa(int(id)) + "/change-password"

	response.HandleRequest(w, r, url)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	url := response.UserServiceRoundRobin.Next().Host + "/api/users/" + strconv.FormatUint(id, 10)
	response.HandleRequest(w, r, url)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	url := response.UserServiceRoundRobin.Next().Host + "/api/users"
	response.HandleRequest(w, r, url)
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	_, err := auth.Authentication(w, r, []string{"user", "admin", "worker"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	url := response.UserServiceRoundRobin.Next().Host + "/api/users" + strconv.FormatUint(id, 10) + "/block"
	response.HandleRequest(w, r, url)
}
