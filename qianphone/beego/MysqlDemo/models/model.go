package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// mysql驱动包
	_ "github.com/go-sql-driver/mysql"

	"MysqlDemo/util"
)

func init() {
	driverName := beego.AppConfig.String("driverName")

	orm.RegisterDriver(driverName, orm.DRMySQL)

	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("post")
	dbname := beego.AppConfig.String("dbname")

	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	err := orm.RegisterDataBase("default", driverName, dbConn)
	if err != nil {
		util.LogError("链接数据库出错！")
		return
	}
	util.LogInfo("数据库链接成功！")
}
