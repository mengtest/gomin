package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gomin/model"
	"gomin/orm"
	"gomin/service"
	"gomin/utils"
)

type MenuController struct {
	beego.Controller
}

func (this *MenuController) MenuList() {
	page, _ := this.GetInt("page", 1)
	limit, _ := this.GetInt("limit", 10)
	count, menus := service.MenuList((page-1)*limit, limit)
	this.Data["json"] = utils.ListSuccess(count, menus)
	this.ServeJSON()
}

func (this *MenuController) MenuAdd() {
	var menu orm.Menu
	var err error
	if json.Unmarshal(this.Ctx.Input.RequestBody, &menu); err == nil {
		logs.Debug("MenuAdd data: ", menu)
		dataId := service.MenuAdd(menu)
		logs.Debug("MenuAdd resp: ", dataId)
		this.Data["json"] = utils.Success(dataId)
	} else {
		logs.Error("MenuAdd error: ", err)
		this.Data["json"] = utils.SYS_PARAM_ERROR
	}
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
