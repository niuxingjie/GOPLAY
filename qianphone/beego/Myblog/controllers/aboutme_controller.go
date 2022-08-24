package controllers

type AboutMeController struct {
	BaseController
}

func (c *AboutMeController) Get() {
	c.Data["wechat"] = "微信：0000"
	c.Data["qq"] = "QQ：000000"
	c.Data["tel"] = "Tel：00000"
	c.TplName = "aboutme.html"
}
