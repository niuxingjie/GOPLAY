package controllers

import (
	"Myblog/models"
	"fmt"
)

//homepage
type HomeController struct {
	BaseController
}

func (this *HomeController) Get() {
	this.Prepare()
	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}
	var article_list []models.Article
	article_list, _ = models.FindArticleWithPage(page)
	this.Data["PageCode"] = 1
	this.Data["HashFooter"] = true
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
	this.Data["Content"] = models.MakeHomeBlocks(article_list, this.IsLogin)
	this.TplName = "home.html"
}
