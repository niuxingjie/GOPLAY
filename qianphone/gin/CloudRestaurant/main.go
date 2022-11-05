package main

import (
	"CloudRestaurant/controller"
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

func main() {

	app_config, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}
	app := gin.Default()
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
}
