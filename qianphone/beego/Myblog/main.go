package main

import (
	"github.com/astaxie/beego"

	_ "Myblog/routers"
	_ "Myblog/utils"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
