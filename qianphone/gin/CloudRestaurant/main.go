package main

import (
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

func main() {

	app_config, err := tool.ParseConfig("./config/app.json")
	if err != nil {
		panic(err)
	}
	app := gin.Default()

	if err := app.Run(app_config.AppHost + ":" + app_config.AppPort); err != nil {
		panic(err)
	}

}
