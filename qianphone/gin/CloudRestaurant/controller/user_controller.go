package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (c UserController) UserRouter(router *gin.RouterGroup) {

	user := User{
		Username: "liming",
		Age:      20,
	}
	router.GET("/info", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 200,
			"msg":  "hello world!",
			"data": user,
		})
	})

	router.GET("/username", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 200,
			"msg":  "hello world!",
			"data": user.Username,
		})
	})

	router.GET("/age", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 200,
			"msg":  "hello world!",
			"data": user.Age,
		})
	})
}

type User struct {
	Username string `json:"name"`
	Age      int    `json:"age"`
}
