package controllers

import (
	"fmt"

	"sample/beeblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["Website"] = "beego.me"        //模板使用
	c.Data["Email"] = "astaxie@gmail.com" //模板使用
	c.TplName = "home.html"
	islogin := checkAccount(c.Ctx)
	fmt.Println("islogin", islogin)
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	cate := c.Input().Get("cate")
	fmt.Println("cate", cate)
	topics, err := models.GetAllTopics(cate, true)
	fmt.Println("topics", topics)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics
	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
		return
	}
	c.Data["Categories"] = categories

}
