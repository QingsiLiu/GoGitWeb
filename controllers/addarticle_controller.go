package controllers

import (
	"GoGitWeb/models"
	"fmt"
	"time"
)

type AddarticleController struct {
	BaseController
}

func (a *AddarticleController) Get() {
	a.TplName = "write_article.html"
}

func (a *AddarticleController) Post() {

	//获取表单中的数据
	title := a.GetString("title")
	tags := a.GetString("tags")
	short := a.GetString("short")
	content := a.GetString("content")
	author := a.GetSession("loginuser")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中
	article := models.Article{0, title, tags, short, content, author.(string), time.Now().Unix()}
	_, err := models.AddArticle(article)

	//返回数据给浏览器
	var response map[string]interface{}

	if err == nil {
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	a.Data["json"] = response
	a.ServeJSON()
}
