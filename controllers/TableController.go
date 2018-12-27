package controllers

import "github.com/astaxie/beego"

type TableController struct {
	beego.Controller
}

func (c *TableController) TableList()  {
	c.Layout= "layout.tpl"
	c.TplName = "table.tpl"
}
