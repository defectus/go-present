package main

import "database/sql"
import "log"
import _ "github.com/lib/pq"

func main() {
	db, _ := sql.Open("postgres", "localhost:5432")
	SQL := `SELECT $1`
	rows, _ := db.Query(SQL, "1")
	for rows.Next() {
		column := sql.NullString{}
		rows.Scan(&column)
		log.Println("column:", column)
	}
	tx, _ := db.Begin()
	res, _ := tx.Exec("INSERT INTO dict (key, value) VALUES (1, 'value')")
	log.Printf("result %#v", res)
	tx.Commit()
}
