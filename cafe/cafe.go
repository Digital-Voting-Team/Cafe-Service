package cafe

import "Cafe-Service/utils"

type Cafe struct {
	Id        int
	CafeName  string `db:"cafe_name"`
	AddressId int    `db:"address"`
	Rating    string
}

func NewCafe(cafeName string, addressId int, rating string) *Cafe {
	return &Cafe{CafeName: cafeName, AddressId: addressId, Rating: rating}
}

func GenerateMockCafe() *Cafe {
	return &Cafe{
		CafeName:  utils.RandStringRunes(20),
		AddressId: 5,
		Rating:    "Looks good",
	}
}
