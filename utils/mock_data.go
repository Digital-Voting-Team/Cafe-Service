package utils

import (
	"Cafe-Service/address"
	"Cafe-Service/cafe"
	"math/rand"
)

func GenerateMockAddress() *address.Address {
	return &address.Address{
		Building:   rand.Intn(100000),
		Street:     RandStringRunes(20),
		City:       RandStringRunes(20),
		District:   RandStringRunes(20),
		Region:     RandStringRunes(20),
		PostalCode: RandStringRunes(5),
	}
}

func GenerateMockCafe() *cafe.Cafe {
	return &cafe.Cafe{
		CafeName:  RandStringRunes(20),
		AddressId: 5,
		Rating:    "Looks good",
	}
}
