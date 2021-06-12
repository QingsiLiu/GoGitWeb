package controllers

import (
	"GoGitWeb/models"
	"GoGitWeb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
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
		r.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		r.ServeJSON()
		return
	}

	//注册用户名和密码
	//存储的密码是md5加密之后的数据，在登录验证的时候，需要将密码经过md5加密之后和数据库中的密码进行判断
	password = utils.MD5(password)
	fmt.Println("md5后的密码为：", password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		fmt.Println(err)
		r.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		r.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	r.ServeJSON()
}
