package models

import (
	"GoGitWeb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title, tags, short, content, author, createtime) values (?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

func QueryArticlesWithCon(sql string) ([]Article, error) {
	sql = "select id,title,tags,short,content,author,createtime from article " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, err
}

func QueryArticlesWithTag(tag string) ([]Article, error) {
	tags := strings.Split(tag, "&")
	sql := "where tags in ("
	for index, tag_now := range tags {
		if index < len(tags)-1 {
			sql = sql + tag_now + ","
		} else if index == len(tags)-1 {
			sql = sql + tag_now + ")"
		}
	}
	return QueryArticlesWithCon(sql)
}

func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticlesWithCon(sql)
}

func FindArticleWithPage(page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryArticleWithPage(page, num)
}
