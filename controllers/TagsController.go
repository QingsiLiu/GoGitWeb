package controllers

import "GoGitWeb/models"

type TagsController struct {
	BaseController
}

func (t *TagsController) Get() {
	tags := models.QueryArticlesWithParam("tags")
	t.Data["Tags"] = models.HandleTagsListData(tags)
	t.TplName = "tags.html"
}
