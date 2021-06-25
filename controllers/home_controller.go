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
	//获取标签以及页码数
	tags := h.GetString("tag")
	fmt.Println("tag:", tags)
	page, _ := h.GetInt("page")
	//初始化一个文章列表
	var artList []models.Article

	if len(tags) > 0 {
		//通过标签来展示页面
		artList, _ = models.QueryArticlesWithTag(tags)
		h.Data["HasFooter"] = false
	} else {
		if page <= 0 {
			page = 1
		}
		artList, _ = models.FindArticleWithPage(page)
		h.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		h.Data["HasFooter"] = true
	}
	fmt.Println("IsLogin:", h.IsLogin, h.Loginuser)
	fmt.Println("文章：", artList)
	h.Data["Content"] = models.MakeHomeBlocks(artList, h.IsLogin)
	h.TplName = "home.html"
}
