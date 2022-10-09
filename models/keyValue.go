package models

import "database/sql"

type KeyValue struct {
	Id    int
	Key   string
	Value string
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

func (KeyValue KeyValue) GetValues() string {
	return "key, value"
}
