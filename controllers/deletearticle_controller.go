package controllers

import (
	"GoGitWeb/models"
	"fmt"
	"log"
)

type DeletearticleController struct {
	BaseController
}

func (d *DeletearticleController) Get() {
	artID, _ := d.GetInt("id")
	fmt.Println("删除的文章id为： ", artID)

	_, err := models.DeleteArticle(artID)
	if err != nil {
		log.Println(err)
	}
	d.Redirect("/", 302)
}
