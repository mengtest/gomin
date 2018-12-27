package controllers

import "github.com/astaxie/beego"


type TestController struct {
	beego.Controller
}


func (c *TestController) TestApp() {
	c.Data["Email"] = "dengjunzhen@weshare.com.cn"
	c.Data["Website"] = "dengjunzhen"
	c.TplName = "index.tpl"
}

