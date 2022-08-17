package controllers

import (
	"github.com/astaxie/beego"
)

type CustomController struct {
	beego.Controller
}

func (this *CustomController) GetUserId() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("变量匹配ID:" + id)
	this.Ctx.ResponseWriter.Write([]byte("ID:" + id))
}

func (this *CustomController) GetUserName() {
	name := this.Ctx.Input.Param(":name")
	beego.Info("变量【0-9】+匹配name:" + name)
	this.Ctx.ResponseWriter.Write([]byte("name:" + name))
}
