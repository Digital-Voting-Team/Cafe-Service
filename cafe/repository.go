package cafe

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `CREATE TABLE IF NOT EXISTS public.cafe
	(
		id integer NOT NULL,
		cafe_name character varying COLLATE pg_catalog."default" NOT NULL,
		address integer NOT NULL,
		rating character varying COLLATE pg_catalog."default" NOT NULL,
		CONSTRAINT cafe_pkey PRIMARY KEY (id),
		CONSTRAINT address_id FOREIGN KEY (address)
			REFERENCES public.addresses (id) MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE NO ACTION
	)
	
	TABLESPACE pg_default;
	
	ALTER TABLE IF EXISTS public.cafe
    OWNER to postgres;`

	queryDeleteTable = `DROP TABLE public.cafe`

	queryInsert = `INSERT INTO public.cafe(
	cafe_name, address, rating)
	VALUES ($1, $2, $3);`

	querySelect = `SELECT * FROM public.cafe;`

	queryUpdate = `UPDATE public.cafe
	SET cafe_name=$2, address=$3, rating=$4
	WHERE id=$1;`

	queryDelete = `DELETE FROM public.cafe
	WHERE id=$1;`

	queryCleanDb = `DELETE FROM public.cafe;`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Insert(cafe *Cafe) (int, error) {
	rows, err := repo.db.Queryx(queryInsert, cafe.CafeName, cafe.AddressId, cafe.Rating)
	defer rows.Close()
	id := -1
	if err != nil {
		return id, err
	}

	rows.Next()
	err = rows.Scan(&id)
	return id, err
}

func (repo *Repository) CreateTable() error {
	_, err := repo.db.Exec(queryCreateTable)
	return err
}

func (repo *Repository) DeleteTable() error {
	_, err := repo.db.Exec(queryDeleteTable)
	return err
}

func (repo *Repository) Select() ([]Cafe, error) {
	rows, err := repo.db.Queryx(querySelect)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	cafe := Cafe{}
	var cafeArray []Cafe
	for rows.Next() {
		err = rows.StructScan(&cafe)
		if err != nil {
			return nil, err
		}
		cafeArray = append(cafeArray, cafe)
	}
	return cafeArray, err
}

func (repo *Repository) Delete(id int) error {
	_, err := repo.db.Exec(queryDelete, id)
	return err
}

func (repo *Repository) Update(id int, cafe *Cafe) error {
	_, err := repo.db.Queryx(queryUpdate, id, cafe.CafeName, cafe.AddressId, cafe.Rating)
	return err
}

func (repo *Repository) Clean() error {
	_, err := repo.db.Exec(queryCleanDb)
	return err
}
