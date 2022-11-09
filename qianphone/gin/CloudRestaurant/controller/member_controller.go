package controller

import (
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (m MemberController) Router(router *gin.RouterGroup) {

	//  http:localhost:8000/memeber/send?phone=12358698745
	router.GET("/send", m.SendCode)

	router.POST("/login", m.SmsCodeLogin)
}

// 发送手机验证码
func (m MemberController) SendCode(context *gin.Context) {

	// 获取参数
	phone, exists := context.GetQuery("phone")
	if !exists {
		tool.Failure(context, "phone参数获取失败")
		return
	}

	// 发送短信验证码
	if ok := new(service.MememberService).SendCode(phone); !ok {
		tool.Failure(context, "发送错误")
		return
	}
	tool.Success(context, "发送错误")
}

// 手机验证码登录
func (m MemberController) SmsCodeLogin(context *gin.Context) {
	var smsLoginParam param.SmsLoginParam

	// 解析json格式参数
	body := context.Request.Body
	err := tool.ParseJson{}.Decode(body, &smsLoginParam)
	if err != nil {
		tool.Failure(context, err.Error())
		return
	}

	// 登录逻辑处理
	member := service.MememberService{}.SmsCodeLogin(smsLoginParam)
	if member != nil {
		tool.Success(context, member)
		return
	}
	tool.Failure(context, "登录错误！")
}
