package controllers

type AboutmeController struct {
	BaseController
}

func (a *AboutmeController) Get() {
	a.Data["wechat"] = "Wechat:lhq8768q"
	a.Data["qq"] = "QQ:1029806879"
	a.Data["tel"] = "Tel:15209880622"
	a.TplName = "aboutme.html"
}
