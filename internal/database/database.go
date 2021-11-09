package database

import (
	"database/sql"
	"fmt"
	"log"
)

type Database struct {
	Host    string
	dbname  string
	user    string
	passw   string
	disable string
}

func Connect() *sql.DB {

	db := Database{
		Host: "",
		dbname: "",
		user: "",
		passw: "",
		disable: "",
	}

	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", db.user, db.dbname, db.passw,
		db.Host, db.disable)
	database, err := sql.Open("postgres", connection)
	if err != nil {
		log.Println("Error: ", err)
	}
	return database
}
