package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化服务引擎
	engine := gin.Default()

	engine.Use(RequestInfo())

	engine.GET("/query", Query)

	// 运行服务
	if err := engine.Run(":8000"); err != nil {
		log.Fatal(err.Error())
		return
	}
}

func Query(context *gin.Context) {
	fmt.Println("函数处理中。。。。。。")
	context.JSON(200, ResponseJson{
		Code:    200,
		Message: "ok",
		Data:    context.FullPath(),
	})
}

// 中间件的定义是：一个gin.HandlerFunc的函数
// 编写中间件就是：如何写一个属于HandlerFunc类型的函数，然后返回
func RequestInfo() gin.HandlerFunc {
	return func(context *gin.Context) {

		// 请求处理函数前执行
		fmt.Println("fullPath:" + context.FullPath())
		fmt.Println("method:" + context.Request.Method)

		fmt.Println("状态码：" + fmt.Sprintf("%d", context.Writer.Status()))
		context.Next()

		// 请求函数处理后执行
		fmt.Println("状态码：" + fmt.Sprintf("%d", context.Writer.Status()))
	}
}

type ResponseJson struct {
	Code    int
	Message string
	Data    interface{}
}
