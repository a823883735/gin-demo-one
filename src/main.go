package main

import (
	//_ "gin-demo-one/src/databases"
	"fmt"
	"gin-demo-one/src/utils"
)

func main() {
	//fmt.Println("Hello World")
	fmt.Println(utils.Md5([]byte("hello")))
}
