package models

import "gin-demo-one/src/utils"

type User struct {
	Id       string `xorm: "id" json:"id"`
	Name     string `xorm: "name" json:"name"`
	Password string `xorm: "password" json:"password"`
}

func (User) TableName() string {
	return "users"
}

func init() {
	utils.SyncTableStruct(User{})
}
