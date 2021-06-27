package controllers

type AlbumController struct {
	BaseController
}

func (a *AlbumController) Get() {
	a.TplName = "album.html"
}
