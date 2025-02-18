package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func initConfig() *mysql.Config {
	return &mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_URL"),
		DBName: os.Getenv("DB_NAME"),
	}
}

func DB() *sql.DB {
	db, err := sql.Open("mysql", initConfig().FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}
