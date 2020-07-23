package models

type User struct {
	Id       string `xorm: "id" json:"id"`
	Name     string `xorm: "name" json:"name"`
	Password string `xorm: "password" json:"password"`
}
