package routers

import (
	"GoGitWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//首页
	beego.Router("/", &controllers.HomeController{})
	//注册
	beego.Router("/register", &controllers.RegisterController{})
	//登录
	beego.Router("/login", &controllers.LoginController{})
	//退出
	beego.Router("/exit", &controllers.ExitController{})
	//写文章
	beego.Router("/article/add", &controllers.AddarticleController{})
	//展示文章内容
	beego.Router("/article/:id", &controllers.ShowArticleController{})
	//删除文章
	beego.Router("/article/delete", &controllers.DeletearticleController{})
	//修改更新文章
	beego.Router("/article/update", &controllers.UpdatearticleController{})
	//根据标签功能展示文章
	beego.Router("/tags", &controllers.TagsController{})
	//相册
	beego.Router("/album", &controllers.AlbumController{})
	//提交相册文件
	beego.Router("/upload", &controllers.UploadController{})
	//关于我
	beego.Router("/aboutme", &controllers.AboutmeController{})
}
