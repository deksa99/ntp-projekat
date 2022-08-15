package handler

import (
	"net/http"
	"os"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {

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
