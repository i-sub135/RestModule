package model

import (
	. "../Struct"
)

func GetData() []RoleUser {
	var roleUsers []RoleUser
	db.Find(&roleUsers)
	return roleUsers
}

func GetDataHistory() []ViewHistory {
	var viewHistory []ViewHistory
	db.Find(&viewHistory)
	return viewHistory
}
