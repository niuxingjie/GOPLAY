package controller

import (
	"CloudRestaurant/service"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (m MemberController) Router(router *gin.RouterGroup) {

	//  http:localhost:8000/memeber/send?phone=12358698745
	router.GET("/send", m.SendCode)
}

func (m MemberController) SendCode(context *gin.Context) {

	// 获取参数
	phone, exists := context.GetQuery("phone")
	if !exists {
		context.JSON(401, map[string]interface{}{
			"code": 40001,
			"msg":  "phone参数获取失败",
			"data": nil,
		})
		return
	}

	// 发送短信验证码
	if ok := new(service.MememberService).SendCode(phone); !ok {
		context.JSON(401, map[string]interface{}{
			"code": 40001,
			"msg":  "发送错误",
			"data": nil,
		})
		return
	}
	context.JSON(200, map[string]interface{}{
		"code": 20001,
		"msg":  "发送成功！",
		"data": nil,
	})
}
