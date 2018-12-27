package controllers

import "github.com/astaxie/beego"

type TemplateController struct {
	beego.Controller
}

func (c *TemplateController) UserListTpl() {
	c.Layout = "layout.tpl"
	c.TplName = "user/userList.tpl"
}

func (c *TemplateController) AddUserTpl() {
	//c.Layout = "layout.tpl"
	c.TplName = "user/addUser.tpl"
}
