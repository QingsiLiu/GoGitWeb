package main

import (
	_ "GoGitWeb/models"
	_ "GoGitWeb/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
