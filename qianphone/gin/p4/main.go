package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// 将map数据以json数据返回
	engine.GET("/hello/json", func(context *gin.Context) {
		fmt.Println("请求路径：", context.FullPath())
		context.JSON(200, map[string]interface{}{
			"code":    1000200,
			"message": "ok",
			"data":    "Hello gin!",
		})
	})

	// 将strcut数据以json格式返回
	engine.GET("/hello/struct/json1", func(context *gin.Context) {
		fmt.Println("请求路径：", context.FullPath())
		context.JSON(200, ResponseJson{
			Code:    1000200,
			Message: "ok",
			Data:    "Hello gin!",
		})
	})

	engine.GET("/hello/struct/json2", func(context *gin.Context) {
		fmt.Println("请求路径：", context.FullPath())
		res := ResponseJson{
			Code:    1000200,
			Message: "ok",
			Data:    "Hello gin!",
		}
		context.JSON(200, &res)
		// context.JSON(200, res) 也可以
	})

	// 配置模板目录
	engine.LoadHTMLGlob("./html/*")
	// 配置静态资源目录(类似django settings中static_url路由前缀和staticsfiles文件存储目录)
	engine.Static("/js", "./static/js")

	// 加载变量及模板渲染（类型django的模板渲染）
	engine.GET("/hello/html", func(context *gin.Context) {
		fullpath := context.FullPath()
		fmt.Println("请求路径：", fullpath)

		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullpath": fullpath,
		})
	})

	if err := engine.Run(":8000"); err != nil {
		log.Fatal(err.Error())
	}
}

type ResponseJson struct {
	Code    int
	Message string
	Data    interface{}
}
