package main

import (
	"gin-demo-one/src/libs"
	"gin-demo-one/src/router"
)

func main() {
	router.Router.Run(":" + libs.Conf.Read("site", "httpport"))
}
