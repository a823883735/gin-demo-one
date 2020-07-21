package main

import (
	"fmt"
	_ "gin-demo-one/src/databases"
	"gin-demo-one/src/libs"
)

func main() {
	//fmt.Println("Hello World")
	var c libs.Configs
	fmt.Println(c)
	var con *libs.Configs
	fmt.Println(con)
}
