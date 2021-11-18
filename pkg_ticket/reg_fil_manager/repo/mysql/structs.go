package mysql

import "github.com/jmoiron/sqlx"

type RegFilRepo struct {
	db *sqlx.DB
}

type dbRegion struct {
	ID   uint64 `db:"region_id"`
	Name string `db:"region"`
}

type dbFilial struct {
	ID       uint64 `db:"filial_id"`
	RegionID uint64 `db:"region_id"`
	Name     string `db:"filial"`
	IP       string `db:"ip"`
}
