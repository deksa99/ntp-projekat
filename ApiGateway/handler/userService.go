package handler

import (
	"ApiGateway/util/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	role := r.URL.Query().Get("role")
	url := response.UserServiceRoundRobin.Next().Host + "/api/users/login?role=" + role

	handleRequest(w, r, url)
}

func Register(w http.ResponseWriter, r *http.Request) {
	url := response.UserServiceRoundRobin.Next().Host + "/api/users/register"

	handleRequest(w, r, url)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	url := response.UserServiceRoundRobin.Next().Host + "/api/users/change-password"
	// TODO authorize

	handleRequest(w, r, url)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	url := response.UserServiceRoundRobin.Next().Host + "/api/users/" + strconv.FormatUint(id, 10)
	handleRequest(w, r, url)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	url := response.UserServiceRoundRobin.Next().Host + "/api/users"
	handleRequest(w, r, url)
}

func BlockUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	// TODO authorize

	url := response.UserServiceRoundRobin.Next().Host + "/api/users" + strconv.FormatUint(id, 10) + "/block"
	handleRequest(w, r, url)
}

func handleRequest(w http.ResponseWriter, r *http.Request, url string) {
	response.CorsResponseHeaders(&w)
	if r.Method == "OPTIONS" {
		return
	}

	req, _ := http.NewRequest(r.Method, url, r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	response.DelegateResponse(res, w)
}
