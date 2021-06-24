package models

import (
	"GoGitWeb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

//文章结构体
type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

//用于添加一篇新的文章，进一步封装
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

//将文章插入数据库里
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title, tags, short, content, author, createtime) values (?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

//将数据库中的文章删除
func DeleteArticle(artid int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", artid)
}

//对数据库中的文章操作
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

//通过文章的id来查询文章
func QueryArticlesWithId(id int) Article {
	row := utils.QueryRowDB("select id, title, tags, short, content, author, createtime from article where id = " + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64 = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

//通过标签来查询文章
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

//通过页码来查询文章
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticlesWithCon(sql)
}

//通过页码来查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryArticleWithPage(page, num)
}

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

//查询文章条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select conut(id) from article")
	nums := 0
	row.Scan(&nums)
	return nums
}

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}