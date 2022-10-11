package repositories

import (
	"FirstGolangProject/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	dbConnection *sql.DB
	Model        models.BaseModel
}

func NewConnection(db *sql.DB, model models.BaseModel) *Connection {
	return &Connection{
		dbConnection: db,
		Model:        model,
	}
}

func (db *Connection) All() []models.BaseModel {
	query, _ := db.dbConnection.Query("select * from " + db.Model.GetTableName())
	defer query.Close()
	var modelsSlice []models.BaseModel
	for query.Next() {
		model := db.Model.NewInstance()
		model.SetValues(query)
		modelsSlice = append(modelsSlice, model)
	}
	fmt.Println(modelsSlice)
	return modelsSlice
}

func (db *Connection) Insert() {
	query, _ := db.dbConnection.Query("insert into " + db.Model.GetTableName() + " (" + db.Model.GetValues() + ") values ('sajjad', 'rabiee')")
	defer query.Close()
}

func (db *Connection) FindById(id int) models.BaseModel {
	query := db.dbConnection.QueryRow("select * from "+db.Model.GetTableName()+" where id = ?", id)
	db.Model.SetValue(query)

	return db.Model
}
