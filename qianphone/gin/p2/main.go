package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		fmt.Println("请求路径：", context.FullPath())
		name := context.DefaultQuery("name", "default_name")
		fmt.Println("请求参数：", name)
		context.Writer.Write([]byte("Hello World," + name))
	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println("请求路径：", context.FullPath())

		username := context.PostForm("username")
		password := context.PostForm("password")

		fmt.Println("请求参数：username=", username)
		fmt.Println("请求参数：password=", password)

		context.Writer.Write([]byte(username + " 正在登录！"))
	})

	if err := engine.Run(":8000"); err != nil {
		log.Fatal(err.Error())
	}
}
