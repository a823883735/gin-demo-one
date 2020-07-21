package databases

import (
	"gin-demo-one/src/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine

const (
	DRIVER_NAME        = "mysql"
	MYSQL_NAME         = "root"
	MYSQL_PWD          = "toor"
	MYSQL_DEV_IP       = "127.0.0.1"
	MYSQL_PRO_IP       = "127.0.0.1"
	MYSQL_PORT         = "3306"
	DATABASE_NAME      = "book"
	CONNECTION_SEETING = "charset=utf8"
	//DATASOURCENAME	= "root:toor@tcp(127.0.0.1:3306)/book?charset=utf8"
)

func init() {
	sqlConnStr := MYSQL_NAME + ":" + MYSQL_PWD
	sqlConnStr += "@tcp(" + MYSQL_DEV_IP + ":" + MYSQL_PORT + ")/"
	sqlConnStr += DATABASE_NAME
	sqlConnStr += "?" + CONNECTION_SEETING
	var err error
	DB, err = xorm.NewEngine(DRIVER_NAME, sqlConnStr)
	utils.ErrorHandle(err, "faild to open mysql connection")

	//err = DB.Sync(new(ObjectType))
	utils.ErrorHandle(err, "faild to sync format table of databases")

	//defer DB.Close()
}
