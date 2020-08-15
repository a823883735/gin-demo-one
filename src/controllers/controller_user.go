package router

import (
	"encoding/base64"
	"gin-demo-one/src/controllers"
	_ "gin-demo-one/src/models"
	"github.com/gin-gonic/gin"
	"github.com/wumansgy/goEncrypt"
	"net/http"
	"strconv"
	"strings"
	"time"
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
			for _, v := range controllers.IGNORE_PATH {
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
			decodePlainText, _ := base64.StdEncoding.DecodeString(token)
			plainText, err := goEncrypt.DesCbcDecrypt(decodePlainText, controllers.PUBLIC_KEY_BYTE_ARRAY)
			plainTextStr := string(plainText)
			index := strings.Index(plainTextStr, "&")
			if err != nil || index < 0 {
				c.JSONP(http.StatusPermanentRedirect, controllers.Result{
					Code: 9999,
					Msg:  "非法token",
				})
				c.Abort()
				return
			}
			if timeStamp, _ := strconv.ParseInt(plainTextStr[index+1:], 10, 64); int(time.Since(time.Unix(timeStamp, 0)).Hours()/24) < 7 {
				c.Set("token", plainTextStr[:index])
				c.Next()
			} else {
				c.JSONP(http.StatusPermanentRedirect, controllers.Result{
					Code: 10001,
					Msg:  "token过期",
				})
				c.Abort()
				return
			}
		}
	}
}

//str, _ := goEncrypt.DesCbcEncrypt([]byte("010f0aea-b5d8-45ef-a27c-148365fc1e53"), []byte(PUBLIC_KEY))
//strText := string(str)
//fmt.Println(strText)

//timeStamp := time.Unix(1597045523, 0)
//fmt.Println(int(time.Since(timeStamp).Hours() / 24) < 7)

//fmt.Println()
//fmt.Println("------------------------------")
//plainText := []byte("010f0aea-b5d8-45ef-a27c-148365fc1e53" + "&" + time.Stamp)
//fmt.Println([]byte(PUBLIC_KEY))
//fmt.Println("明文：", string(plainText))
//cryptText, _ := goEncrypt.DesCbcEncrypt(plainText, PUBLIC_KEY_BYTE_ARRAY)
//fmt.Println("密文：", base64.StdEncoding.EncodeToString(cryptText))
//
//str, _ := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(cryptText))
//fmt.Println("密文：", base64.StdEncoding.EncodeToString(str))
//newPlainText, _ := goEncrypt.DesCbcDecrypt(str, []byte("@t0!K1nl"))
//fmt.Println("明文：", string(newPlainText))
//fmt.Println("------------------------------")
