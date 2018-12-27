package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"gomin/orm"
	"gomin/service"
	"gomin/utils"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) AddUser() {
	var user orm.User
	var err error
	if json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		fmt.Println("userName: ", user.Name)
		objectId := service.AddUser(user)
		this.Data["json"] = utils.Success(objectId)
	} else {
		this.Data["json"] = utils.SysError()
	}
	this.ServeJSON()
}

func (this *UserController) UserList() {
	page, _ := this.GetInt("page", 1)
	limit, _ := this.GetInt("limit", 10)
	count, userList := orm.UserList((page-1)*limit, limit)
	this.Data["json"] = utils.ListSuccess(count, userList)
	this.ServeJSON()
}
