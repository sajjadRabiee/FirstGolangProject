package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func Connection() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/first_golang_project")
	if err != nil {
		log.Print(err.Error())
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	//query, _ := db.Query("insert into key_value_table (`key`, `value`) values ('sajjad', 'rabiee')")
	//defer query.Close()
	defer db.Close()
}
