package controllers

type HomeController struct {
	BaseController
}

// Get 博客首页地址，localhost：8080
func (h *HomeController) Get() {

	h.TplName = "home.html"
}
