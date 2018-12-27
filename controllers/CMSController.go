package controllers

import "github.com/astaxie/beego"

type CMSController struct {
	beego.Controller
}


// @router /test/router [get]
func (c *CMSController) TestAnno() {
	c.Data["Email"] = "dengjunzhen@weshare.com.cnqqqqqqqqqqqqqqqqqqqqqbbbbbbbbbbbbbbb"
	c.Data["Website"] = "dengjunzhenqqqqqqqqqqqqqqqq"
	c.TplName = "index.tpl"
}
