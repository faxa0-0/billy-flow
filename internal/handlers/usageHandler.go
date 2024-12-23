package handlers

import "net/http"

func GetUsage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Usage"))
}

func UpdateUsageByUserID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Usage By User ID"))
}
