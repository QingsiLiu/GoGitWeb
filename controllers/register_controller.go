package controllers

import (
	"GoGitWeb/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

//处理注册
func (r *RegisterController) Post() {
	//获取表单信息
	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")
	fmt.Println(username, password, repassword)
	logs.Info(username, password, repassword)

	//注册之前先判断该用户是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id", id)
	if id > 0 {
		r.Data["json"] = 1
	}

	//注册用户

}
