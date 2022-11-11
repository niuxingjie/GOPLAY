package controllers

import (
	"github.com/astaxie/beego"
)

type RgController struct {
	beego.Controller
}

func (this *RgController) Get() {
	id := this.Ctx.Input.Param(":id")
	beego.Info("变量匹配ID:" + id)
	this.Ctx.ResponseWriter.Write([]byte("ID:" + id))

	name := this.Ctx.Input.Param(":name")
	beego.Info("变量【0-9】+匹配name:" + name)
	this.Ctx.ResponseWriter.Write([]byte("name:" + name))

	age := this.Ctx.Input.Param(":age")
	beego.Info("变量int匹配age:" + age)
	this.Ctx.ResponseWriter.Write([]byte("age:" + age))

	info := this.Ctx.Input.Param(":info")
	beego.Info("变量string匹配age:" + info)
	this.Ctx.ResponseWriter.Write([]byte("info:" + info))
}
