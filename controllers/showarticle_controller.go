package controllers

import (
	"GoGitWeb/models"
	"fmt"
)

type ShowArticleController struct {
	BaseController
}

func (s *ShowArticleController) Get() {
	artId, _ := s.GetInt(":id")
	fmt.Println("文章的id为： ", artId)

	art := models.QueryArticlesWithId(artId)
	s.Data["Title"] = art.Title
	s.Data["Content"] = art.Content
	s.TplName = "show_article.html"
}
