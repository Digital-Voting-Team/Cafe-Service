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

func (p *cafesQ) New() data.CafesQ {
	return NewCafesQ(p.db)
}

func (p *cafesQ) Get() (*data.Cafe, error) {
	var result data.Cafe
	err := p.db.Get(&result, p.sql)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (p *cafesQ) Select() ([]data.Cafe, error) {
	var result []data.Cafe
	err := p.db.Select(&result, p.sql)
	return result, err
}

func (p *cafesQ) Update(cafe data.Cafe) (data.Cafe, error) {
	var result data.Cafe
	clauses := structs.Map(cafe)
	clauses["cafe_name"] = cafe.CafeName
	clauses["rating"] = cafe.Rating
	clauses["address_id"] = cafe.AddressId

	err := p.db.Get(&result, p.sqlUpdate.SetMap(clauses))
	return result, err
}

func (p *cafesQ) Transaction(fn func(q data.CafesQ) error) error {
	return p.db.Transaction(func() error {
		return fn(p)
	})
}

func (p *cafesQ) Insert(cafe data.Cafe) (data.Cafe, error) {
	clauses := structs.Map(cafe)
	clauses["cafe_name"] = cafe.CafeName
	clauses["rating"] = cafe.Rating
	clauses["address_id"] = cafe.AddressId

	var result data.Cafe
	stmt := sq.Insert(cafesTableName).SetMap(clauses).Suffix("returning *")
	err := p.db.Get(&result, stmt)

	return result, err
}

func (p *cafesQ) Delete(id int64) error {
	stmt := sq.Delete(cafesTableName).Where(sq.Eq{"id": id})
	err := p.db.Exec(stmt)
	return err
}

func (p *cafesQ) Page(pageParams pgdb.OffsetPageParams) data.CafesQ {
	p.sql = pageParams.ApplyTo(p.sql, "id")
	return p
}

func (p *cafesQ) FilterById(ids ...int64) data.CafesQ {
	p.sql = p.sql.Where(sq.Eq{"id": ids})
	p.sqlUpdate = p.sqlUpdate.Where(sq.Eq{"id": ids})
	return p
}

func (p *cafesQ) FilterByNames(names ...string) data.CafesQ {
	p.sql = p.sql.Where(sq.Eq{"cafe_name": names})
	return p
}

func (p *cafesQ) FilterByRatingFrom(ratings ...float64) data.CafesQ {
	stmt := sq.GtOrEq{"cafes.rating": ratings}
	p.sql = p.sql.Where(stmt)
	return p
}

func (p *cafesQ) FilterByRatingTo(ratings ...float64) data.CafesQ {
	stmt := sq.LtOrEq{"cafes.rating": ratings}
	p.sql = p.sql.Where(stmt)
	return p
}

func (p *cafesQ) JoinAddress() data.CafesQ {
	stmt := fmt.Sprintf("%s as cafes on public.addresses.id = cafes.address_id",
		cafesTableName)
	p.sql = p.sql.Join(stmt)
	return p
}
