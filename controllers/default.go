package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//func Test(c *MainController) {
//	c.Data["Email"] = "dengjunzhen@weshare.com.cn"
//	c.Data["Website"] = "dengjunzhen"
//	c.TplName = "index.tpl"
//}
