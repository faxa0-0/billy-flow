package handlers

import "net/http"

func GenerateInvoice(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Generate Invoice"))
}

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Invoice"))
}
