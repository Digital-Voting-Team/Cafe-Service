package address

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	queryCreateTable = `CREATE TABLE IF NOT EXISTS public.addresses
	(
		id integer NOT NULL,
		building_number integer NOT NULL,
		street character varying COLLATE pg_catalog."default" NOT NULL,
		city character varying COLLATE pg_catalog."default" NOT NULL,
		district character varying COLLATE pg_catalog."default" NOT NULL,
		region character varying COLLATE pg_catalog."default" NOT NULL,
		postal_code character varying COLLATE pg_catalog."default" NOT NULL,
		CONSTRAINT addresses_pkey PRIMARY KEY (id)
	)
	
	TABLESPACE pg_default;
	
	ALTER TABLE IF EXISTS public.addresses
    OWNER to postgres;`

	queryDeleteTable = `DROP TABLE public.addresses`

	queryInsert = `INSERT INTO public.addresses(
	building, street, city, district, region, postal_code)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	querySelect = `SELECT * FROM public.addresses;`

	queryUpdate = `UPDATE public.addresses
	SET building=$2, street=$3, city=$4, district=$5, region=$6, postal_code=$7
	WHERE id=$1;`

	queryDelete = `DELETE FROM public.addresses
	WHERE id=$1;`
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Insert(addr *Address) (int, error) {
	rows, err := repo.db.Queryx(queryInsert, addr.Building, addr.Street, addr.City, addr.District, addr.Region, addr.PostalCode)
	defer rows.Close()
	id := -1
	if err != nil {
		return id, err
	}

	rows.Next()
	err = rows.Scan(&id)
	return id, nil
}

func (repo *Repository) CreateTable() error {
	_, err := repo.db.Exec(queryCreateTable)
	return err
}

func (repo *Repository) DeleteTable() error {
	_, err := repo.db.Exec(queryDeleteTable)
	return err
}

func (repo *Repository) Select() ([]Address, error) {
	rows, err := repo.db.Queryx(querySelect)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	addr := Address{}
	addressArray := []Address{}
	for rows.Next() {
		err = rows.StructScan(&addr)
		if err != nil {
			return nil, err
		}
		addressArray = append(addressArray, addr)
	}
	return addressArray, err
}

func (repo *Repository) Delete(id int) error {
	_, err := repo.db.Exec(queryDelete, id)
	return err
}

func (repo *Repository) Update(id int, addr *Address) error {
	_, err := repo.db.Queryx(queryUpdate, id, addr.Building, addr.Street, addr.City, addr.District, addr.Region, addr.PostalCode)
	return err
}
