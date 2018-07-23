package controllers

import (
	"fmt"
	"sample/beeblog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	c.TplName = "topic.html"
	topics, err := models.GetAllTopics("", false)
	fmt.Println("topics", topics)
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topic"] = topics

}

//func (c *TopicController) Add() { //当在连接里面输入add的时候，只能路由会自动匹配到Add()方法。
//	//	c.Ctx.WriteString("add") //会在客户端输出日志。
//}
func (c *TopicController) Post() {
	var err error
	//所有提交的都会到这个里面
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	tid := c.Input().Get("tid")
	category := c.Input().Get("category")

	err = models.AddCategory(category)
	if err != nil {
		fmt.Println("添加文章插入分类报错")
		return
	}

	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)
	} else {
		err = models.ModifyTopic(tid, title, category, content)
	}

	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)
}
func (c *TopicController) Add() {
	c.TplName = "topic_add.html"
}
func (c *TopicController) View() {
	c.TplName = "topic_view.html"
	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(c.Input().Get("tid")) //表示获取？后面的第一个参数c.Ctx.Input.Param("0")
	//	fmt.Println("文章id", topic, c.Ctx.Input.Param("0"))
	if err != nil {
		c.Redirect("/topic", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Tid"] = tid

	replies, err1 := models.GetAllReplies(tid)
	if err1 != nil {
		beego.Error(err1)
		return
	}
	c.Data["Replies"] = replies
	c.Data["IsLogin"] = checkAccount(c.Ctx) //在模板里面有进行判断，判断是不是登录了，不登陆不能进行删除
}
func (c *TopicController) Modify() {
	c.TplName = "topic_modify.html"
	tid := c.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/topic", 302)
		return
	}
	c.Data["Topic"] = topic
	c.Data["Tid"] = tid

}
func (c *TopicController) Delete() {
	if !checkAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopic(c.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic", 302)

}
