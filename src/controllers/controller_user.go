package controllers

import (
	"gin-demo-one/src/databases"
	"gin-demo-one/src/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func GetUsers(c *gin.Context) {
	table := databases.DB.SQL("select * from users")

	table = table.Limit(0, 4)
	//var u models.User
	var list []models.User
	defer table.Close()
	if err := table.Find(&list); err == nil {
		c.JSONP(http.StatusOK, Result{
			Code: 200, Data: Page{
				List: list,
			},
		})
		count := len(list)
		table.Limit(0, 4).Find(&list)
	} else {
		c.JSONP(http.StatusOK, Result{
			Code: 100,
			Msg:  "操作失败",
		})
	}

	return
	//rows, err := databases.DB.Rows(models.User{})
	//if err != nil {
	//	utils.ErrorHandle(err, "error")
	//	c.JSONP(http.StatusOK, Result{Code: 100, Msg: "操作失败"})
	//} else {
	//	var list []models.User
	//	defer rows.Close()
	//	var user models.User
	//	for rows.Next() {
	//		err = rows.Scan(&user)
	//		if err == nil {
	//			list = append(list, user)
	//		}
	//	}
	//	c.JSONP(http.StatusOK, Result{Code: 200, Data: list})
	//}
}

func AddUser(c *gin.Context) {
	user := new(models.User)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("pwd")
	var err error
	u1 := uuid.Must(uuid.NewV4(), err)
	user.Id = u1.String()
	if affected, err := databases.DB.Insert(user); err != nil || affected == 0 {
		c.JSONP(http.StatusOK, Result{Code: 100, Msg: "操作失败"})
	} else {
		c.JSONP(http.StatusOK, Result{Code: 200, Data: user})
	}
}

func DeleteUser(c *gin.Context) {

}

func UpdataUser(c *gin.Context) {

}

func SelectUser(c *gin.Context) {

}

func Login(c *gin.Context) {

}
