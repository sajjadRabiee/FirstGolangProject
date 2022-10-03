package repositories

import (
	"FirstGolangProject/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	connection *sql.DB
	model      models.BaseModel
}

func NewConnection(db *sql.DB, model models.BaseModel) *Connection {
	return &Connection{
		connection: db,
		model:      model,
	}
}

func (db *Connection) All() []models.BaseModel {
	query, _ := db.connection.Query("select * from " + db.model.GetTableName())
	defer query.Close()
	var modelsSlice []models.BaseModel
	for query.Next() {
		model := db.model.NewInstance()
		model.SetValues(query)
		modelsSlice = append(modelsSlice, model)
	}
	fmt.Println(modelsSlice)
	return modelsSlice
}

func (db *Connection) Insert() {
	query, _ := db.connection.Query("insert into " + db.model.GetTableName() + " (`key`, `value`) values ('sajjad', 'rabiee')")
	defer query.Close()
}

func (db *Connection) FindById(id int) models.BaseModel {
	query := db.connection.QueryRow("select * from "+db.model.GetTableName()+" where id = ?", id)
	db.model.SetValue(query)

	return db.model
}
