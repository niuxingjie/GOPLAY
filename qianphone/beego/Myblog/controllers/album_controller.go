package controllers

import (
	"Myblog/models"

	"github.com/opentracing/opentracing-go/log"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		log.Error(err)
	}
	this.Data["Album"] = albums
	this.Data["IsLogin"] = this.IsLogin
	this.TplName = "album.html"
}
