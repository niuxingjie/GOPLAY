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

	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}
	var article_list []models.Article

	// 查询文章数据
	article_list, _ = models.FindArticleWithPage(page)

	// 组合页码数据
	this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
	this.Data["HasFooter"] = true
	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)

	// 文章数据渲染html源码
	this.Data["Content"] = models.MakeHomeBlocks(article_list, this.IsLogin)
	this.TplName = "home.html"
}
