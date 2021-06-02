package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

//固定路由的get方法
//func (this *MainController) Get() {
//	beego.Info("固定路由的get方法")
//	this.Ctx.Output.Body([]byte("固定路由get请求"))
//}

//固定路由的post方法
//func (this *MainController) Post() {
//	beego.Info("固定路由的post方法")
//	this.Ctx.Output.Body([]byte("固定路由post请求"))
//}

/*
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
*/
