package controllers

import (
	"GoGitWeb/models"
	"fmt"
)

type UpdatearticleController struct {
	BaseController
}

//访问update时触发get方法，通过tplname返回客户端页面
func (u *UpdatearticleController) Get() {
	artId, _ := u.GetInt("id")
	fmt.Println("修改的文章id为：", artId)

	art := models.QueryArticlesWithId(artId)

	u.Data["Title"] = art.Title
	u.Data["Tags"] = art.Tags
	u.Data["Short"] = art.Short
	u.Data["Content"] = art.Content
	u.Data["Id"] = art.Id
	u.TplName = "write_article.html"
}

//修改文章
func (u *UpdatearticleController) Post() {
	artId, _ := u.GetInt("id")
	fmt.Println("当前修改的文章id为：", artId)

	title := u.GetString("title")
	tags := u.GetString("tags")
	short := u.GetString("short")
	content := u.GetString("content")
	author := u.GetSession("loginuser")
	fmt.Println("修改文章的是：", author)

	//实例化文章结构体，更新数据库
	article := models.Article{artId, title, tags, short, content, author.(string), 0}
	_, err := models.UpdateArticle(article)

	if err != nil {
		u.Data["json"] = map[string]interface{}{"code": 0, "message": "修改失败"}
	} else {
		u.Data["json"] = map[string]interface{}{"code": 1, "message": "修改成功"}
	}

	u.ServeJSON()
}
