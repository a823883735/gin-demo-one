package models

import (
	"gin-demo-one/src/databases"
	"gin-demo-one/src/utils"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}

func init() {
	err := databases.DB.Sync(new(User))
	utils.ErrorHandle(err, "faild to sync format table of databases")
}
