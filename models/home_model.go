package models

import (
	"GoGitWeb/utils"
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

// TagLink 标签结构体
type TagLink struct {
	TagName string
	TagUrl  string
}

// HomeFooterPageCode 分页结构体
type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

// MakeHomeBlocks ----------首页显示内容-----------
func MakeHomeBlocks(article []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range article {
		//将数据库model转换为首页模板所需要的model
		//先实例化一个首页文章模块并赋值
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.Short = art.Short
		homeParam.CreateTime = utils.SwitchTimeStampToData(art.Createtime)
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		fmt.Println("首页展示文章模块字符串内容", t)
		buffer := bytes.Buffer{}
		err := t.Execute(&buffer, homeParam)
		if err != nil {
			fmt.Println("err: ", err)
			return ""
		}
		htmlHome += buffer.String()
	}
	return template.HTML(htmlHome)
}

func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsParam := strings.Split(tags, "&")
	for _, tag := range tagsParam {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}

// ConfigHomeFooterPageCode ----------分页/翻页------------  page是当前页数
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	nums := GetArticleRowNum()
	//读取配置的每页条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	//计算出总页数
	fmt.Println("文章的总数量为：", nums)
	sumPageNum := (nums-1)/pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, sumPageNum)

	//如果当前页数小于等于1，那么不存在上一页
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}
	//如果当前页数大于等于总页数，那么不存在下一页
	if page >= sumPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}

	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)

	return pageCode
}
