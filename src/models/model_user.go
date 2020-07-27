package models

import (
	"gin-demo-one/src/databases"
	"gin-demo-one/src/utils"
	_ "github.com/go-xorm/xorm"
)

type User struct {
	Id       string `json:"id" xorm:"notnull pk varchar(50) 'id'"`
	Name     string `json:"name" xorm:"varchar(12)"`
	Phone    string `json:"phone" xorm:"notnull varchar(11)"`
	Password string `json:"password" xorm:"notnull varchar(100)"`
}

func (User) TableName() string {
	return "users"
}

func init() {
	err := databases.DB.Sync(new(User))
	utils.ErrorHandle(err, "faild to sync format table of databases")
}
