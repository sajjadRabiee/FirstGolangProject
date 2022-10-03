package repositories

import (
	"FirstGolangProject/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	connection *sql.DB
}

func NewConnection(db *sql.DB) *Connection {
	return &Connection{
		connection: db,
	}
}

func (db *Connection) All() []models.KeyValue {
	query, _ := db.connection.Query("select * from key_value_table")
	defer query.Close()
	var model models.KeyValue
	var keyValues []models.KeyValue
	for query.Next() {
		query.Scan(&model.Id, &model.Key, &model.Value)
		keyValues = append(keyValues, model)
	}
	fmt.Println(keyValues)
	return keyValues
}

func (db *Connection) Insert() {
	query, _ := db.connection.Query("insert into key_value_table (`key`, `value`) values ('sajjad', 'rabiee')")
	defer query.Close()
}

func (db *Connection) FindById(id int) models.KeyValue {
	query := db.connection.QueryRow("select * from key_value_table where id = ?", id)
	var model models.KeyValue
	query.Scan(&model.Id, &model.Key, &model.Value)

	return model
}
