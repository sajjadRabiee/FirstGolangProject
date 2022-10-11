package models

import (
	"database/sql"
	"fmt"
)

type KeyValue struct {
	Id    int    `json:"id,omitempty"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (keyValue *KeyValue) SetValue(row *sql.Row) {
	row.Scan(&keyValue.Id, &keyValue.Key, &keyValue.Value)
}

func (keyValue *KeyValue) SetValues(row *sql.Rows) {
	row.Scan(&keyValue.Id, &keyValue.Key, &keyValue.Value)
}

func (keyValue KeyValue) GetTableName() string {
	return "key_values"
}

func (keyValue KeyValue) NewInstance() BaseModel {
	return &KeyValue{}
}

func (keyValue KeyValue) GetTableTitles() string {
	return "`key`, `value`"
}

func (keyValue *KeyValue) GetValues() string {
	fmt.Println(keyValue)
	return "'" + keyValue.Key + "', '" + keyValue.Value + "'"
}
