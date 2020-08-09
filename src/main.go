package main

import (
	"gin-demo-one/src/libs"
	"gin-demo-one/src/router"
)
func main() {
	router.Router.Run(":" + libs.Conf.Read("site", "httpport"))
}

//func main(){
//	app := gin.Default()
//	app.Use(Cors())
//	app.NoRoute(func(c *gin.Context) {
//		a := c.Request.RequestURI
//		data, _ := regexp.Match("^/img/(.+)$", []byte(a))
//		if data {
//			reg, _ := regexp.Compile("^/img/(.+)$")
//			str := reg.ReplaceAllString(a, "$1")
//			str = str[strings.Index(str, ",")+1:]
//			dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
//			c.Set("Content-Type", "image/jpg")
//			io.Copy(c.Writer, dec)
//			//ddd, _ := base64.StdEncoding.DecodeString(str)
//			//生成图片并保存
//			//err := ioutil.WriteFile("test.png", ddd, 7777)
//		}
//	})
//	//app.POST("/", func(c *gin.Context) {
//	//	c.JSONP(http.StatusOK, gin.H{
//	//		"total": 28,
//	//		"rows": []gin.H{
//	//			{"productid": "FI-SW-01", "unitcost": 10.00, "status": "P", "listprice": 16.50, "attr1": "Large", "itemid": "EST-1"},
//	//			{"productid": "K9-DL-01", "unitcost": 12.00, "status": "P", "listprice": 18.50, "attr1": "Spotted Adult Female", "itemid": "EST-10"},
//	//			{"productid": "RP-SN-01", "unitcost": 12.00, "status": "P", "listprice": 18.50, "attr1": "Venomless", "itemid": "EST-11"},
//	//			{"productid": "RP-SN-01", "unitcost": 12.00, "status": "P", "listprice": 18.50, "attr1": "Rattleless", "itemid": "EST-12"},
//	//			{"productid": "RP-LI-02", "unitcost": 12.00, "status": "P", "listprice": 18.50, "attr1": "Green Adult", "itemid": "EST-13"},
//	//			{"productid": "FL-DSH-01", "unitcost": 12.00, "status": "P", "listprice": 58.50, "attr1": "Tailless", "itemid": "EST-14"},
//	//			{"productid": "FL-DSH-01", "unitcost": 12.00, "status": "P", "listprice": 23.50, "attr1": "With tail", "itemid": "EST-15"},
//	//			{"productid": "FL-DLH-02", "unitcost": 12.00, "status": "P", "listprice": 93.50, "attr1": "Adult Female", "itemid": "EST-16"},
//	//			{"productid": "FL-DLH-02", "unitcost": 12.00, "status": "P", "listprice": 93.50, "attr1": "Adult Male", "itemid": "EST-17"},
//	//			{"productid": "AV-CB-01", "unitcost": 92.00, "status": "P", "listprice": 193.50, "attr1": "Adult Male", "itemid": "EST-18"},
//	//		},
//	//	})
//	//})
//
//	app.GET("/img/:str", func(c *gin.Context) {
//		str := c.Param("str")
//		dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(str))
//		c.Header("Content-Type", "image/png")
//		io.Copy(c.Writer, dec)
//	})
//	app.Run()
//}

//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		method := c.Request.Method
//
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
//		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
//		c.Header("Access-Control-Allow-Credentials", "true")
//
//		//放行所有OPTIONS方法
//		if method == "OPTIONS" {
//			c.AbortWithStatus(http.StatusNoContent)
//		}
//		// 处理请求
//		c.Next()
//	}
//}
