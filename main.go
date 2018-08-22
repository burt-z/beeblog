package main

import (
	"sample/beeblog/controllers"
	_ "sample/beeblog/routers"

	"sample/beeblog/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, false)
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	//	beego.AutoRouter(&controllers.ReplyController{}) //使用beego的只能路由，
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.AutoRouter(&controllers.TopicController{}) //使用beego的只能路由，
	beego.Run()
	//提交测试1
}
