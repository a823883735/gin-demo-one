package databases

import (
	"gin-demo-one/src/libs"
	"gin-demo-one/src/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strconv"
)

var DB *xorm.Engine

func init() {
	var err error
	userName := libs.Conf.Read("mysql", "username")
	password := libs.Conf.Read("mysql", "password")
	dataname := libs.Conf.Read("mysql", "dataname")
	host := libs.Conf.Read("mysql", "host")
	port := libs.Conf.Read("mysql", "port")
	connStr := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataname + "?charset=utf8&parseTime=true"

	DB, err = xorm.NewEngine("mysql", connStr)
	DB.ShowSQL(true)
	utils.ErrorHandle(err, "faild to open mysql connection")

	maxIdleConns, err := strconv.ParseInt(libs.Conf.Read("mysql", "maxIdleConn"), 10, 0)
	utils.ErrorHandle(err, "this maxIdleConns expeced.")
	DB.SetMaxIdleConns(int(maxIdleConns))

	maxOpenConns, err := strconv.ParseInt(libs.Conf.Read("mysql", "maxOpenConn"), 10, 0)
	utils.ErrorHandle(err, "this maxOpenConns expeced.")
	DB.SetMaxIdleConns(int(maxOpenConns))

	//err = DB.Sync(new(ObjectType))
	//utils.ErrorHandle(err, "faild to sync format table of databases")

	//defer DB.Close()
}
