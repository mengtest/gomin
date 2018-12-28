package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gomin/model"
	"gomin/orm"
	"gomin/service"
	"gomin/utils"
)

type MenuController struct {
	beego.Controller
}

func (this *MenuController) MenuList() {
	offset, _ := this.GetInt("offset", 0)
	limit, _ := this.GetInt("limit", 10)
	count, menus := orm.MenuList(offset, limit)
	this.Data["json"] = utils.ListSuccess(count, menus)
	this.ServeJSON()
}

func (this *MenuController) MenuDelete() {
	var req model.MenuDeleteReq
	var err error
	if json.Unmarshal(this.Ctx.Input.RequestBody, &req); err == nil {
		count, err := service.MeneDelete(req)
		if err == nil {
			this.Data["json"] = utils.Success(count);
		} else {
			this.Data["json"] = utils.SysError();
		}
	} else {
		this.Data["json"] = utils.SysError();
	}
	this.ServeJSON()
}
