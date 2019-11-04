package model

import (
	Struct "../Struct"
)

// GetData -- ORM Get Data User
func GetData() []Struct.RoleUser {
	var roleUsers []Struct.RoleUser
	db.Find(&roleUsers)
	return roleUsers
}

// GetDataHistory -- ORM Get data History
func GetDataHistory() []Struct.ViewHistory {
	var viewHistory []Struct.ViewHistory
	db.Find(&viewHistory)
	return viewHistory
}
