package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=0.0.0.0 port=5422 user=rinha password=rinha dbname=rinha_db sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db, err
}
