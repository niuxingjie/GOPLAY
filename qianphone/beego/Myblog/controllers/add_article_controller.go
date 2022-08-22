package controllers

import (
	"Myblog/models"
	"fmt"
	"time"
)

type AddArticleController struct {
	BaseController
}

func (this *AddArticleController) Get() {
	this.Prepare()
	if this.IsLogin {
		this.TplName = "write_article.html"
		return
	} else {
		this.Redirect("/login/", 302)
	}
}

func (this *AddArticleController) Post() {
	this.Prepare()
	if this.IsLogin {
		title := this.GetString("title")
		tags := this.GetString("tags")
		short := this.GetString("short")
		content := this.GetString("content")
		fmt.Printf("title:%s,tags:%s\n", title, tags)

		// TODO:位置参数与关键字参数如何实现？
		// 0这里是给id占位
		article := models.Article{0, title, tags, short, content, "author", time.Now().Unix()}
		_, err := models.AddArticle(&article)

		var response map[string]interface{}
		if err != nil {
			fmt.Println("err:", err)
			response = map[string]interface{}{"code": 0, "message": "failed"}
		} else {
			response = map[string]interface{}{"code": 1, "message": "ok"}
		}
		this.Data["json"] = response
		this.ServeJSON()
	} else {
		this.Redirect("/login/", 302)
	}
}
