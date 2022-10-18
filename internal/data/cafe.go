package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
)

type CafesQ interface {
	New() CafesQ

	Get() (*Cafe, error)
	Select() ([]Cafe, error)

	Transaction(fn func(q CafesQ) error) error

	Insert(cafe Cafe) (Cafe, error)
	Update(cafe Cafe) (Cafe, error)
	Delete(id int64) error

	Page(pageParams pgdb.OffsetPageParams) CafesQ

	FilterById(ids ...int64) CafesQ
	FilterByNames(names ...string) CafesQ
	FilterByRatingFrom(ratings ...float64) CafesQ
	FilterByRatingTo(ratings ...float64) CafesQ

	JoinAddress() CafesQ
}

type Cafe struct {
	Id        int64   `db:"id" structs:"-"`
	CafeName  string  `db:"cafe_name" structs:"cafe_name"`
	AddressId int64   `db:"address_id" structs:"address_id"`
	Rating    *string `db:"rating" structs:"rating"`
}
