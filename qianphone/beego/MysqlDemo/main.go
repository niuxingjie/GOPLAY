package main

import (
	_ "MysqlDemo/models"
	_ "MysqlDemo/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
