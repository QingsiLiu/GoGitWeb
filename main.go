package main

import (
	_ "GoGitWeb/models"
	_ "GoGitWeb/routers"
	"GoGitWeb/utils"
	"github.com/astaxie/beego"
)

func main() {
	utils.InitMysql()
	beego.Run()
}
