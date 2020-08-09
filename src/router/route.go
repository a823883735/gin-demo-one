package router

import (
	"crypto/rsa"
	"gin-demo-one/src/controllers"
	_ "gin-demo-one/src/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/timestamp"
	"io"
	"math/rand"
	"net/http"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(IsLogin())
	Router.Use(Cors())
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

var IGNORE_PATH = [...]string{"/user/login", "/user/register", "/user/find"}
var PUBLIC_KEY = "token"

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			token = c.PostForm("token")
		} else {
			token = c.Query("token")
		}
		if token == "" {
			requestPath := c.FullPath()
			for _, v := range IGNORE_PATH {
				if requestPath == v {
					c.Next()
					return
				}
			}
			c.JSONP(http.StatusPermanentRedirect, controllers.Result{
				Code: 10000,
				Msg:  "未登录",
			})
			c.Abort()
		} else {
			//timeStatmp, err := rsa.DecryptPKCS1v15(io.Reader{}, byte[](PUBLIC_KEY),timestamp.Timestamp{} )
			c.Next()
		}
	}
}
