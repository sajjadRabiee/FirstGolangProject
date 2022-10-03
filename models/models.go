package models

import "database/sql"

type BaseModel interface {
	SetValue(row *sql.Row)
	SetValues(row *sql.Rows)
	GetTableName() string
	NewInstance() BaseModel
}
