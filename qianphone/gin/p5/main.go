package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	userRouter := engine.Group("/user")

	userRouter.POST("/register", register)
	userRouter.POST("/login", login)
	userRouter.POST("/delete", delete)

	if err := engine.Run(":8000"); err != nil {
		log.Fatal("err:" + err.Error())
	}
}

func register(context *gin.Context) {
	fmt.Println("fullPath:" + context.FullPath())

	var user User
	err := context.ShouldBind(&user) // 绑定不会报错
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(user)
	context.Writer.Write([]byte(user.Username + "' 's register succeed!"))
}

func login(context *gin.Context) {
	fmt.Println("fullPath:" + context.FullPath())

	var user User
	err := context.ShouldBind(&user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(user)
	context.Writer.Write([]byte(user.Username + "' login succeed!"))
}

func delete(context *gin.Context) {
	fmt.Println("fullPath:" + context.FullPath())

	var user User
	err := context.ShouldBind(&user)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	fmt.Println(user)
	context.Writer.Write([]byte(user.Username + "' delete succeed!"))
}

type User struct {
	Username string `form:"name"`
	Password string `form:"password"`
}
