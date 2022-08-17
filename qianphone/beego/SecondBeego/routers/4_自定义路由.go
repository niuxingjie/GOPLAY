package routers

import (
	"SecondBeego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/getUserId/:id", &controllers.CustomController{}, "GET:GetUserId")

	beego.Router("/getUserName/:name([0-9]+)", &controllers.CustomController{}, "GET:GetUserName")
}
