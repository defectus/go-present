package main

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type OurDataRow struct {
	Key   int            `db:"key"`
	Value sql.NullString `db:"value"`
}

func ReadWithMapping(db *sql.DB) *OurDataRow {
	ourDataRow := &OurDataRow{}
	dbx := sqlx.NewDb(db, "postgres")
	dbx.Select(ReadWithMapping, "SELECT * FROM dict")
	return ourDataRow
}
