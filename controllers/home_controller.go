package controllers

import (
	"GoGitWeb/models"
	"fmt"
)

type HomeController struct {
	BaseController
}

// Get 博客首页地址，localhost：8080
func (h *HomeController) Get() {

	tag := h.GetString("tag")
	fmt.Println("tag: ", tag)
	page, _ := h.GetInt("page")
	var artlist []models.Article

	if len(tag) > 0 {
		//按照指定的标签进行搜索
		artlist, _ = models.QueryArticlesWithTag(tag)
		h.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		//设置分页
		artlist, _ = models.FindArticleWithPage(page)
		h.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		h.Data["HasFooter"] = true
	}

	fmt.Println("IsLogin:", h.IsLogin, h.Loginuser)
	h.Data["Content"] = models.MakeHomeBlocks(artlist, h.IsLogin)
	h.TplName = "home.html"
}
