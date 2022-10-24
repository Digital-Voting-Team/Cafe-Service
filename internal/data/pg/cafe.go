package pg

import (
	"Cafe-Service/internal/data"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const cafesTableName = "public.cafes"

func NewCafesQ(db *pgdb.DB) data.CafesQ {
	return &cafesQ{
		db:        db.Clone(),
		sql:       sq.Select("cafes.*").From(cafesTableName),
		sqlUpdate: sq.Update(cafesTableName).Suffix("returning *"),
	}
}

type cafesQ struct {
	db        *pgdb.DB
	sql       sq.SelectBuilder
	sqlUpdate sq.UpdateBuilder
}

func (q *cafesQ) New() data.CafesQ {
	return NewCafesQ(q.db)
}

func (q *cafesQ) Get() (*data.Cafe, error) {
	var result data.Cafe
	err := q.db.Get(&result, q.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *cafesQ) Select() ([]data.Cafe, error) {
	var result []data.Cafe
	err := q.db.Select(&result, q.sql)
	return result, err
}

func (q *cafesQ) Update(cafe data.Cafe) (data.Cafe, error) {
	var result data.Cafe
	clauses := structs.Map(cafe)
	clauses["cafe_name"] = cafe.CafeName
	clauses["rating"] = cafe.Rating
	clauses["address_id"] = cafe.AddressId

	err := q.db.Get(&result, q.sqlUpdate.SetMap(clauses))
	return result, err
}

func (q *cafesQ) Transaction(fn func(q data.CafesQ) error) error {
	return q.db.Transaction(func() error {
		return fn(q)
	})
}

func (q *cafesQ) Insert(cafe data.Cafe) (data.Cafe, error) {
	clauses := structs.Map(cafe)
	clauses["cafe_name"] = cafe.CafeName
	clauses["rating"] = cafe.Rating
	clauses["address_id"] = cafe.AddressId

	var result data.Cafe
	stmt := sq.Insert(cafesTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)

	return result, err
}

func (q *cafesQ) Delete(id int64) error {
	stmt := sq.Delete(cafesTableName).Where(sq.Eq{"id": id})
	err := q.db.Exec(stmt)
	return err
}

func (q *cafesQ) Page(pageParams pgdb.OffsetPageParams) data.CafesQ {
	q.sql = pageParams.ApplyTo(q.sql, "id")
	return q
}

func (q *cafesQ) FilterById(ids ...int64) data.CafesQ {
	q.sql = q.sql.Where(sq.Eq{"id": ids})
	q.sqlUpdate = q.sqlUpdate.Where(sq.Eq{"id": ids})
	return q
}

func (q *cafesQ) FilterByNames(names ...string) data.CafesQ {
	q.sql = q.sql.Where(sq.Eq{"cafe_name": names})
	return q
}

func (q *cafesQ) FilterByRatingFrom(ratings ...float64) data.CafesQ {
	stmt := sq.GtOrEq{"rating": ratings}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *cafesQ) FilterByRatingTo(ratings ...float64) data.CafesQ {
	stmt := sq.LtOrEq{"rating": ratings}
	q.sql = q.sql.Where(stmt)
	return q
}

func (q *cafesQ) FilterByAddressId(ids ...int64) data.CafesQ {
	q.sql = q.sql.Where(sq.Eq{"address_id": ids})
	return q
}

func (q *cafesQ) JoinAddress() data.CafesQ {
	stmt := fmt.Sprintf("%s as cafes on public.addresses.id = cafes.address_id",
		cafesTableName)
	q.sql = q.sql.Join(stmt)
	return q
}
