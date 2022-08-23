package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

func (this *BaseController) Prepare() {
	// TODO：这个貌似是beego.Controller的接口，会自动执行？
	// TODO: 后端保存session是如何实现的？貌似数据库里没有？http请求是无状态的，怎样保证，同一客户端请求发过来，记录登录状态？
	// TODO: 重启后session就失效了？
	loginuser := this.GetSession("loginuser")
	fmt.Println("loginuser--->", loginuser)
	if loginuser != nil {
		this.IsLogin = true
		this.Loginuser = loginuser
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}
