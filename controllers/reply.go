package controllers

import (
	"sample/beeblog/models"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Add() {
	tid := c.Input().Get("tid") //提交的表单里面有三个参数
	nickname := c.Input().Get("nickname")
	content := c.Input().Get("content")
	err := models.AddReply(tid, nickname, content)
	if err != nil {
		beego.Error(err)
		return
	}
	c.Redirect("/topic/view?tid="+tid, 302)
}
func (c *ReplyController) Delete() {
	if !checkAccount(c.Ctx) { //判断是不是已经登录，不登陆不显示
		return
	}
	rid := c.Input().Get("rid")
	tid := c.Input().Get("tid") //在评论里面获取tid，需要模板里面对tid进行操作{{$tid := Topic.Tid}}
	err := models.DeleteReply(rid)
	if err != nil {
		beego.Error(err)

	}
	c.Redirect("/topic/view?tid="+tid, 302)
}
