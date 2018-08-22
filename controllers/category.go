package controllers

import (
	"fmt"
	"sample/beeblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	fmt.Println("op属性", op)
	switch op {
	case "add":
		//name=123&op=add  输入框里面的值
		name := c.Input().Get("name")
		fmt.Println("name", name)
		if (len(name)) == 0 {
			fmt.Println("name 为空")
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			fmt.Println("category 没有查到")
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return

	}
	c.TplName = "category.html"
	c.Data["IsCategory"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	//	第一次提交。

}
