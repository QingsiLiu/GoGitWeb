package routers

//beego通过正则表达式来解析http请求
func init() {
	//*全匹配
	//beego.Router("/*", &controllers.RegController{})

	//:id变量匹配
	//beego.Router("/getUser/:id", &controllers.RegController{})

	////正则表达式自定义匹配
	//  /getUser/davie
	//beego.Router("/getUser/:name([0-9]+)", &controllers.RegController{})

	//*.* 匹配
	//beego.Router("/upload/*.*", &controllers.RegController{})

	//int类型匹配
	//beego.Router("/getUserInfo/:id:int", &controllers.RegController{})

	//string类型匹配
	//beego.Router("/getUserInfo/:username:string", &controllers.RegController{})
}
