package models

import "database/sql"

type Hobby struct {
	Id    int
	Name  string
	Value string
}

func (hobby *Hobby) SetValue(row *sql.Row) {
	row.Scan(&hobby.Id, &hobby.Name, &hobby.Value)
}

func (hobby *Hobby) SetValues(row *sql.Rows) {
	row.Scan(&hobby.Id, &hobby.Name, &hobby.Value)
}

func (hobby Hobby) GetTableName() string {
	return "hobbies"
}

func (hobby Hobby) NewInstance() BaseModel {
	return &Hobby{}
}

func (hobby Hobby) GetValues() string {
	return "name, value"
}
