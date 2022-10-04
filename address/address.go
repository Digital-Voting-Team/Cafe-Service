package address

import (
	"Cafe-Service/utils"
	"math/rand"
)

type Address struct {
	Id         int
	Building   int
	Street     string
	City       string
	District   string
	Region     string
	PostalCode string `db:"postal_code"`
}

func NewAddress(building int, street string, city string, district string, region string, postalCode string) *Address {
	return &Address{Building: building, Street: street, City: city, District: district, Region: region, PostalCode: postalCode}
}

func GenerateMockAddress() *Address {
	return &Address{
		Building:   rand.Intn(100000),
		Street:     utils.RandStringRunes(20),
		City:       utils.RandStringRunes(20),
		District:   utils.RandStringRunes(20),
		Region:     utils.RandStringRunes(20),
		PostalCode: utils.RandStringRunes(5),
	}
}
