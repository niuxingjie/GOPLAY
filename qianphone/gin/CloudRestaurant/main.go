package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

func main() {

	// 读取配置文件
	app_config, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}

	// 初始化数据库引擎
	_, err = tool.OrmEngine(app_config)
	if err != nil {
		panic(err)
	}

	// 初始化服务引擎
	app := gin.Default()
	app.Use(tool.Cors())
	registerRouter(app)
	if err := app.Run(app_config.AppHost + ":" + app_config.AppPort); err != nil {
		panic(err)
	}

}

func registerRouter(engine *gin.Engine) {

	// user
	new(controller.UserController).UserRouter(engine.Group("/user"))

	// Member
	new(controller.MemberController).Router(engine.Group("/member"))

	// captcha
	new(controller.CaptchaController).Router(engine.Group("/captcha"))
}
