package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Host    string
	Port 	int
	dbname  string
	user    string
	passw   string
	disable string
}

func Connect() *sql.DB {

	db := Database{
		Host:    "127.0.0.1",
		dbname:  "izanami",
		user:    "root",
		passw:   "root",
		disable: "disable",
	}

	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", db.user, db.dbname, db.passw,
		db.Host, db.disable)
	fmt.Println(connection)
	database, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println("Error: ", err)
	}
	return database
}
