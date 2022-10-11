package models

import (
	"database/sql"
	"fmt"
)

type Hobby struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Value string `json:"value"`
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

func (hobby Hobby) GetTableTitles() string {
	return "`name`, `value`"
}

func (hobby *Hobby) GetValues() string {
	fmt.Println(hobby)
	return "'" + hobby.Name + "', '" + hobby.Value + "'"
}
