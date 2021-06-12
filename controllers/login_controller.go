package controllers

import (
	"GoGitWeb/models"
	"GoGitWeb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}

// Post 登录功能控制
func (l *LoginController) Post() {

	//获取表单信息
	username := l.GetString("username")
	password := l.GetString("password")
	fmt.Println("用户名为：", username, " 密码为：", password)
	logs.Info(username, password)

	//将密码进行加密转换
	password = utils.MD5(password)
	id := models.QueryUserWithParam(username, password)
	fmt.Println("id:", id)
	if id > 0 {
		/*
			设置了session后悔将数据处理设置到cookie，然后再浏览器进行网络请求的时候回自动带上cookie
			因为我们可以通过获取这个cookie来判断用户是谁，这里我们使用的是session的方式进行设置
		*/
		l.SetSession("loginuser", username)
		l.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		l.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败"}
	}
	l.ServeJSON()
}
