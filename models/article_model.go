package models

import (
	"GoGitWeb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

// Article 文章结构体
type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

// AddArticle 用于添加一篇新的文章，进一步封装
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

//将文章插入数据库里
func insertArticle(article Article) (int64, error) {
	return utils.ModifyDB("insert into article(title, tags, short, content, author, createtime) values (?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

// DeleteArticle 将数据库中的文章删除
func DeleteArticle(artid int) (int64, error) {
	return utils.ModifyDB("delete from article where id=?", artid)
}

// UpdateArticle 更新数据库中的文章
func UpdateArticle(article Article) (int64, error) {
	return utils.ModifyDB("update article set title=?, tags=?, short=?, content=?, author=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Id)
}

// QueryArticlesWithCon 对数据库中的文章操作
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

// QueryArticlesWithId 通过文章的id来查询文章
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

// QueryArticlesWithParam 根据列名字段来查询文章（此处用作标签文章的展示），返回一个标签列表
func QueryArticlesWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}
	var paramlist []string
	for rows.Next() {
		tmp := ""
		rows.Scan(&tmp)
		paramlist = append(paramlist, tmp)
	}
	return paramlist
}

// QueryArticlesWithTag 通过标签来查询文章
func QueryArticlesWithTag(tag string) ([]Article, error) {
	sql := "where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	fmt.Println("通过标签查询文章sql：", sql)
	return QueryArticlesWithCon(sql)
}

// QueryArticleWithPage 通过页码来查询文章
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticlesWithCon(sql)
}

// FindArticleWithPage 通过页码来查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	num, _ := beego.AppConfig.Int("articleListPageNum")
	page--
	fmt.Println("---------->page", page)
	return QueryArticleWithPage(page, num)
}

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

// QueryArticleRowNum 查询文章条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	nums := 0
	row.Scan(&nums)
	return nums
}

// GetArticleRowNum 只有首次获取行数的时候采取统计表里的行数
func GetArticleRowNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}
