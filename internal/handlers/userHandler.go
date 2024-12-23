package handlers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User By ID"))
}
