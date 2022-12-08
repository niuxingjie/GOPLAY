package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"WebChatRoom/controller"
	"WebChatRoom/tool"
)

func main() {

	engine := gin.Default()

	// 启用多模板及静态文件服务
	engine.Static("/static", "./static")
	engine.HTMLRender = tool.LoadTemplates("./templates")

	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	engine.POST("/login", controller.Login)
	engine.GET("/room", controller.Room)
	engine.GET("/get_chat_list", controller.GetChatList)
	engine.POST("/post_message", controller.PostMessage)

	if err := engine.Run("0.0.0.0:8000"); err != nil {
		log.Fatal(err.Error())
	}
}
