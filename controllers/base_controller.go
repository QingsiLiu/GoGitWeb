package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin   bool
	Loginuser interface{}
}

//判断是否登录
/*
	这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，
    用户可以重写这个函数实现类似用户验证之类。
*/
func (b *BaseController) Prepare() {
	loginuser := b.GetSession("loginuser")
	fmt.Println("loginuser------>", loginuser)
	if loginuser != nil {
		b.IsLogin = true
		b.Loginuser = loginuser
	} else {
		b.IsLogin = false
	}
	b.Data["IsLogin"] = b.IsLogin
}
