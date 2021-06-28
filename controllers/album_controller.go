package controllers

import (
	"GoGitWeb/models"
	"github.com/astaxie/beego/logs"
)

type AlbumController struct {
	BaseController
}

func (a *AlbumController) Get() {
	albums, err := models.FindAlbum()
	if err != nil {
		logs.Error(err)
	}
	a.Data["Album"] = albums
	a.TplName = "album.html"
}
