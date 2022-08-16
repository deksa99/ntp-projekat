package handler

import (
	"ApiGateway/util/response"
	"net/http"
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

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {

}

func BlockUser(w http.ResponseWriter, r *http.Request) {

}

func handleRequest(w http.ResponseWriter, r *http.Request, url string) {
	response.CorsResponseHeaders(&w)
	if r.Method == "OPTIONS" {
		return
	}

	req, _ := http.NewRequest(http.MethodPost, url, r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	response.DelegateResponse(res, w)
}
