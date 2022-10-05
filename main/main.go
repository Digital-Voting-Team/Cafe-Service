package main

import (
	"github.com/jmoiron/sqlx"
	"log"
)

func Connect(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func main() {
	connStr := "user=postgres dbname=Cafe sslmode=disable password=password"
	db, err := Connect(connStr)

	if err != nil {
		log.Fatal(err)
	}

	AddressesSimulation(db)
	CafeSimulation(db)
}
