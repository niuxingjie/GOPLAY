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

	var articleList []models.Article

	tag := this.GetString("tag")
	if len(tag) > 0 {
		articleList, _ = models.QueryArticleWithTag(tag)
		this.Data["HasFooter"] = false
	} else {
		page, _ := this.GetInt("page")
		if page <= 0 {
			page = 1
		}
		// 查询文章数据
		articleList, _ = models.FindArticleWithPage(page)

		// 组合页码数据
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}

	fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)

	// 文章数据渲染html源码
	this.Data["Content"] = models.MakeHomeBlocks(articleList, this.IsLogin)
	this.TplName = "home.html"
}
