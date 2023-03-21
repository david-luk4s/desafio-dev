package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func ConnectionDB() *sql.DB {
	DB, err = sql.Open("postgres", "host=db port=5432 dbname=desafiodev user=postgres password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		log.Panic(err.Error())
	}

	return DB
}
