package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func StartDB(host string, port int, user string, password string, dbName string) {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	var err error
	db, err = sql.Open("postgres", info)
	if err != nil {
		fmt.Println(db)
		log.Fatal("Koneksi ke database gagal!")
	}
}

func GetDB() *sql.DB {
	return db
}
