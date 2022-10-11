package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/first_golang_project")
	if err != nil {
		log.Print(err.Error())
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	return db
}
