package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"

	"Myblog/models"
	"Myblog/utils"
)

type MainController struct {
	beego.Controller
}

// TODO:这里并没有发现又返回结果的代码啊？
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)

	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)

	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在了！"}
		this.ServeJSON()
		return
	}

	password = utils.MD5(password)
	fmt.Println("md5后:", password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败！"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册成功！"}
	}
	this.ServeJSON()
}

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")

	id := models.QueryUserWithParam(username, utils.MD5(password))
	if id > 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录成功！"}
		fmt.Println("id:", id)
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败！"}
	}
	this.ServeJSON()
}
