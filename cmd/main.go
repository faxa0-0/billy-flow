package main

import (
	"log"
	"net/http"

	"github.com/faxa0-0/billy-flow/internal/config"
	"github.com/faxa0-0/billy-flow/internal/db"
	"github.com/faxa0-0/billy-flow/internal/handlers"
)

func main() {
	cfg := config.LoadEnv()

	err := db.Initialize(cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/login", handlers.Login) //all

	mux.HandleFunc("POST /users", handlers.CreateUser)      // admin
	mux.HandleFunc("GET /users", handlers.GetUsers)         // admin
	mux.HandleFunc("GET /users/{id}", handlers.GetUserByID) //all

	mux.HandleFunc("GET /usage/{id}", handlers.GetUsage)            // all there's; admin -all's
	mux.HandleFunc("PUT /usage/{id}", handlers.UpdateUsageByUserID) //admin

	mux.HandleFunc("POST /invoice", handlers.GenerateInvoice) // admin
	mux.HandleFunc("GET /invoice/{id}", handlers.GetInvoice)  //all

	log.Print("Server started at http://" + cfg.FormatAddr())
	log.Fatal(http.ListenAndServe(cfg.FormatAddr(), mux))
}
