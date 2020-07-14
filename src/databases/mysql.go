package databases

import (
	"gin-demo-one/src/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine

const (
	DRIVER_NAME    = "mysql"
	DATASOURCENAME = "root:toor@tcp(127.0.0.1:3306)/book?charset=utf8"
)

func init() {
	var err error
	DB, err = xorm.NewEngine(DRIVER_NAME, DATASOURCENAME)
	utils.ErrorHandle(err, "faild to open mysql connection")

	err = DB.Sync(new(ObjectType))
	utils.ErrorHandle(err, "faild to sync format table of databases")

	defer DB.Close()
}
