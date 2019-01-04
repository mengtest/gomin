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
	c.TplName = "user/addUser.tpl"
}

func (c *TemplateController) EditUserTpl() {
	c.TplName = "user/editUser.tpl"
	c.Data["userId"] = c.Ctx.Input.Param(":id")
}

func (c *TemplateController) MenuListTpl()  {
	c.Layout = "layout.tpl"
	c.TplName = "menu/menuList.tpl"
}

func (c *TemplateController) MenuAddTpl() {
	c.TplName = "menu/addMenu.tpl"
}
