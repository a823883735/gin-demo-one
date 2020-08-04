package router

import (
	"gin-demo-one/src/controllers"
	_ "gin-demo-one/src/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"net/http"
	"strconv"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(Cors())
	Router.GET("/mail", func(c *gin.Context) {
		mailConn := map[string]string{
			"user": "516948767@qq.com",
			"pass": "sprzkxacqdhabhcc",
			"host": "smtp.qq.com",
			"port": "465", //587
		}
		port, _ := strconv.Atoi(mailConn["port"])

		m := gomail.NewMessage()

		m.SetHeader("From", m.FormatAddress(mailConn["user"], "测试"))

		m.SetHeader("To", "823883735@qq.com")
		m.SetHeader("Subject", "邮箱测试")
		var text = c.Query("text")
		m.SetBody("text/html", text)

		d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

		if err := d.DialAndSend(m); err == nil {
			c.JSONP(http.StatusOK, controllers.Result{
				Code: 200,
				Data: "success",
			})
		} else {
			c.JSONP(http.StatusOK, controllers.Result{
				Code: 100,
				Msg:  "faild",
			})
		}
	})
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
