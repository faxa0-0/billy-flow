package types

import "time"

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Login          string `json:"login"`
	HashedPassword string `json:"hashed_pass"`
	Role           string `json:"role"`
	Active         bool   `json:"active"`
}

type Usage struct {
	CustomerID int    `json:"customer_id"`
	Youtube    uint64 `json:"youtube"` //kb
	Netflix    uint64 `json:"netflix"`
	Spotify    uint64 `json:"spotify"`
	Base       uint64 `json:"basic"`
}

type Invoices struct {
	ID          int       `json:"id"`
	CustomerID  int       `json:"customer_id"`
	Fee         uint64    `json:"fee"`
	GeneratedAt time.Time `json:"generated_at"`
}
