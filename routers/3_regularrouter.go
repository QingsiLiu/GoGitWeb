package routers

import (
	"GoGitWeb/controllers"
	"github.com/astaxie/beego"
)

//beego通过正则表达式来解析http请求
func init() {
	//*全匹配
	beego.Router("/*", &controllers.MainController{})

	//:id变量匹配
	beego.Router("/getUser/:id", &controllers.MainController{})

	////正则表达式自定义匹配
	//  /getUser/davie
	beego.Router("/getUser/:name([0-9]+)", &controllers.MainController{})
}
