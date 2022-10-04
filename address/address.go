package address

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
