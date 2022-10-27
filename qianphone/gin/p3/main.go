package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化服务引擎
	engine := gin.Default()

	engine.GET("/info", func(context *gin.Context) {

		fmt.Println("请求路径：", context.FullPath())

		// get 参数实体绑定 ShouldBindQuery
		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println("student:", student)
		context.Writer.Write([]byte(student.Name))
	})

	engine.POST("/register", func(context *gin.Context) {

		fmt.Println("请求路径：", context.FullPath())

		// post 参数实体绑定 ShouldBind
		var register Register
		err := context.ShouldBind(&register)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println("register:", register)
		context.Writer.Write([]byte(register.Phone))
	})

	engine.POST("/register/json", func(context *gin.Context) {

		fmt.Println("请求路径：", context.FullPath())

		// post 参数实体绑定 BindJSON
		var register Register
		err := context.BindJSON(&register)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println("register:", register)
		context.Writer.Write([]byte(register.Phone))
	})

	// 运行服务
	if err := engine.Run(":8000"); err != nil {
		log.Fatal(err.Error())
		return
	}
}

// 结构体的参数名使用大驼峰才能导出
type Student struct {
	Name  string `form:"name"`
	Class string `form:"class"`
}

type Register struct {
	Username string `form:"name"` // 这里可以把name映射到Username
	Phone    string `form:"phone"`
	Password string `form:"password"`
}
