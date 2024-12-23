package main

import (
	"log"
	"net/http"

	"github.com/faxa0-0/billy-flow/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/login", handlers.Login) //all

	mux.HandleFunc("POST /users", handlers.CreateUser)      // admin
	mux.HandleFunc("GET /users", handlers.GetUsers)         // admin
	mux.HandleFunc("GET /users/{id}", handlers.GetUserByID) //all

	mux.HandleFunc("GET /usage/{id}", handlers.GetUsage)            // all there's; admin -all's
	mux.HandleFunc("PUT /usage/{id}", handlers.UpdateUsageByUserID) //admin

	mux.HandleFunc("POST /invoice", handlers.GenerateInvoice) // admin
	mux.HandleFunc("GET /invoice/{id}", handlers.GetInvoice)  //all

	log.Fatal(http.ListenAndServe(":8080", mux))
}
