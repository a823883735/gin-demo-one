package controllers

import (
	"fmt"
	"gin-demo-one/src/databases"
	"gin-demo-one/src/models"
	"gin-demo-one/src/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func GetUsers(c *gin.Context) {
	list := []models.User{}
	page := NewPage(c.Query("pageNum"), c.Query("pageSize"))
	//result, err := GetListSplitPage("select * from users", &list, c.Query("pageNum"), c.Query("pageSize"))
	table := databases.DB.SQL("select sql_calc_found_rows * from users")
	defer table.Close()
	if err := table.Find(&list); err == nil {
		table.Close()
		total := len(list)
		q := fmt.Sprint("select id, name, phone,'' password from users limit ", (page.PageNum-1)*page.PageSize, ",", page.PageSize)
		list = []models.User{}
		databases.DB.SQL(q).Find(&list)
		page.GetListSplitPage(list, total, len(list))
		c.JSONP(http.StatusOK, Result{
			Code: 200,
			Data: page,
		})
	} else {
		c.JSONP(http.StatusOK, Result{
			Code: 100,
			Msg:  "操作失败",
		})
	}
	return
}

func AddUser(c *gin.Context) {
	var err error
	user := new(models.User)
	user.Phone = c.PostForm("phone")
	user.Password = c.PostForm("pwd")
	user.Password, err = utils.GeneratePassword(user.Password)
	if err == nil {
		u1 := uuid.Must(uuid.NewV4(), err)
		user.Id = u1.String()
		if _, err := databases.DB.Insert(user); err == nil {
			c.JSONP(http.StatusOK, Result{Code: 200, Data: user})
			return
		}
	}
	c.JSONP(http.StatusOK, Result{Code: 100, Msg: "操作失败"})
}

func DeleteUser(c *gin.Context) {
	user := new(models.User)
	user.Id = c.Param("id")
	if _, err := databases.DB.Id(user.Id).Delete(&user); err == nil {
		c.JSONP(http.StatusOK, Result{Code: 200, Data: user})
	} else {
		c.JSONP(http.StatusOK, Result{Code: 100, Msg: "操作失败"})
	}
}

func UpdataUser(c *gin.Context) {
	var err error
	user := new(models.User)
	user.Id = c.PostForm("id")
	user.Phone = c.PostForm("phone")
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("pwd")
	user.Password, err = utils.GeneratePassword(user.Password)
	if err == nil {
		if _, err = databases.DB.Id(user.Id).Update(user); err == nil {
			c.JSONP(http.StatusOK, Result{Code: 200, Data: user})
			return
		}
	}
	c.JSONP(http.StatusOK, Result{Code: 100, Msg: "操作失败"})
}

func SelectUser(c *gin.Context) {
	user := new(models.User)
	user.Id = c.Query("id")
	if _, err := databases.DB.Id(user.Id).Get(user); err == nil {
		c.JSONP(http.StatusOK, Result{Code: 200, Data: user})
	} else {
		c.JSONP(http.StatusOK, Result{Code: 100, Msg: "操作失败"})
	}
}

func Login(c *gin.Context) {
	var err error
	user := new(models.User)
	user.Phone = c.PostForm("phone")
	pwd := c.PostForm("pwd")
	_, err = databases.DB.Where("phone=?", user.Phone).Get(user)
	if err == nil {
		if ok, err := utils.VaildataPassword(pwd, user.Password); err == nil {
			if ok {
				c.Redirect(http.StatusMovedPermanently, "/index.html")
			} else {
				c.JSONP(http.StatusOK, Result{Code: 100, Msg: "用户名与密码不匹配"})
			}
			return
		}
	}
	c.JSONP(http.StatusOK, Result{Code: 500, Msg: "服务器异常，请联系管理员"})
}
