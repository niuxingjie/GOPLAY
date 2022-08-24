package routers

import (
	"Myblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register/", &controllers.RegisterController{})
	beego.Router("/login/", &controllers.LoginController{})
	beego.Router("/exit/", &controllers.ExitController{})

	// 文章管理相关接口
	beego.Router("/article/add", &controllers.AddArticleController{})
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	beego.Router("/article/update", &controllers.UpdateArticleController{})
	beego.Router("/article/delete", &controllers.DeleteArticleController{})
}
