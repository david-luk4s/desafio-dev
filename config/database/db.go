package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB     *sql.DB
	DBTest *sql.DB
	err    error
)

func ConnectionDB() *sql.DB {
	DB, err = sql.Open("postgres", "host=db port=5432 dbname=desafiodev user=postgres password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		log.Panic(err.Error())
	}

	return DB
}

func ConnectionDBTest() *sql.DB {
	//remove test.db case exist
	os.Remove("./test.db")

	DBTest, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
		log.Panic(err.Error())
	}

	AutoMigrate()
	return DBTest
}
