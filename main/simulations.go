package main

import (
	"Cafe-Service/address"
	"Cafe-Service/cafe"
	"Cafe-Service/utils"
	"github.com/jmoiron/sqlx"
	"log"
)

func AddressesSimulation(db *sqlx.DB) {
	addressRepo := address.NewRepository(db)
	newAddress := utils.GenerateMockAddress()

	id, err := addressRepo.Insert(newAddress)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("id of the added address: %d", id)

	addrArr, err := addressRepo.Select()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\naddresses: %+v", addrArr)
}

func CafeSimulation(db *sqlx.DB) {
	cafeRepo := cafe.NewRepository(db)
	newCafe := utils.GenerateMockCafe()

	id, err := cafeRepo.Insert(newCafe)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("id of the added cafe: %d", id)

	cafeArr, err := cafeRepo.Select()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\naddresses: %+v", cafeArr)
}
