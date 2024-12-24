package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Initialize(dsn string) error {
	DB, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("could not connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected!")

	if err := createTables(DB); err != nil {
		return fmt.Errorf("could not create tables: %v", err)
	}

	log.Println("Successfully connected to the database and ensured tables exist.")
	return nil
}

func createTables(db *sql.DB) error {
	usersTableQuery := `
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		login VARCHAR(100) UNIQUE NOT NULL,
		hashed_pass VARCHAR(255) NOT NULL,
		role VARCHAR(50) NOT NULL,
		active BOOLEAN DEFAULT TRUE
	);`

	_, err := db.Exec(usersTableQuery)
	if err != nil {
		return fmt.Errorf("could not create users table: %v", err)
	}

	usageTableQuery := `
	CREATE TABLE usage (
		id INT PRIMARY KEY,
		customer_id INT NOT NULL,
		youtube BIGINT NOT NULL,
		netflix BIGINT NOT NULL,
		spotify BIGINT NOT NULL,
		basic BIGINT NOT NULL,
		FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err = db.Exec(usageTableQuery)
	if err != nil {
		return fmt.Errorf("could not create usage table: %v", err)
	}

	invoicesTableQuery := `
	CREATE TABLE invoices (
		id SERIAL PRIMARY KEY,
		customer_id INT NOT NULL,
		fee BIGINT NOT NULL,
		generated_at TIMESTAMP NOT NULL,
		FOREIGN KEY (customer_id) REFERENCES users(id) ON DELETE CASCADE
	);`

	_, err = db.Exec(invoicesTableQuery)
	if err != nil {
		return fmt.Errorf("could not create invoices table: %v", err)
	}

	return nil
}
