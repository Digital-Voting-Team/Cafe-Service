package utils

import (
	"Cafe-Service/address"
	"Cafe-Service/cafe"
	"math/rand"
)

func GenerateMockAddress() *address.Address {
	return &address.Address{
		Building:   rand.Intn(100),
		Street:     RandStringRunes(2),
		City:       RandStringRunes(2),
		District:   RandStringRunes(2),
		Region:     RandStringRunes(2),
		PostalCode: RandStringRunes(5),
	}
}

func GenerateMockCafe() *cafe.Cafe {
	return &cafe.Cafe{
		CafeName:  RandStringRunes(2),
		AddressId: 1,
		Rating:    "Looks good",
	}
}
