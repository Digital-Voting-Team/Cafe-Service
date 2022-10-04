package cafe

type Cafe struct {
	Id        int
	CafeName  string `db:"cafe_name"`
	AddressId int    `db:"address"`
	Rating    string
}

func NewCafe(cafeName string, addressId int, rating string) *Cafe {
	return &Cafe{CafeName: cafeName, AddressId: addressId, Rating: rating}
}
