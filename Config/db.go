package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(){
	db, err := sql.Open("mysql", "test:123@/go.test")//menggunakan db mysql dengan nama configurasi database go.test
	//mengembalikan dua return, yang pertama adalah koneksinya dan error
	if err != nil {
		panic(err)
	}

	log.Println("Database connected")
	DB = db
}