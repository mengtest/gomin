package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gomin/model"
	"gomin/orm"
	"gomin/service"
	"gomin/utils"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) AddUser() {
	var user orm.User
	var err error
	if json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		logs.Debug("userName: ", user.Name)
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

func (this *UserController) DeleteUser() {
	var req model.DeleteUserReq
	var err error
	if json.Unmarshal(this.Ctx.Input.RequestBody, &req); err == nil {
		logs.Debug("DeleteUser ids: ", req)
		count, err := service.DeleteUser(req.UserIds)
		if err != nil {
			logs.Error("DeleteUser error: ", err)
			this.Data["json"] = utils.SysError()
		} else {
			this.Data["json"] = utils.Success(count)
		}
	} else {
		this.Data["json"] = utils.SysError()
	}
	this.ServeJSON()
}

func (this *UserController) UpdateUserStatusMulti() {
	var req model.UpdateUserStatusReq
	var err error
	if json.Unmarshal(this.Ctx.Input.RequestBody, &req); err == nil {
		logs.Debug("DisableUser ids: ", req)
		count, err := service.UpdateUserStatusMulti(req.UserIds, req.TargetStatus)
		if err != nil {
			logs.Error("DisableUser error: ", err)
			this.Data["json"] = utils.SysError()
		} else {
			this.Data["json"] = utils.Success(count)
		}
	} else {
		this.Data["json"] = utils.SysError()
	}
	this.ServeJSON()
}

func (this *UserController) QueryOneUser() {
	userId, e:= strconv.Atoi(this.Ctx.Input.Param(":id"))
	logs.Debug("QueryOneUser userId: ", userId)
	if e != nil {
		this.Data["json"] = utils.SYS_PARAM_ERROR
	} else {
		user := service.QueryOneUser(userId)
		this.Data["json"] = utils.Success(user)
	}
	this.ServeJSON()
}

func (this *UserController) UpdateUser()  {
	var user orm.User
	var err error
	if json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		logs.Debug("UpdateUser userId: ", user.Id)
		result, err := service.UpdateUser(user)
		if err == nil {
			this.Data["json"] = utils.Success(result)
		} else{
			this.Data["json"] = utils.SysError()
		}
	} else {
		this.Data["json"] = utils.SysError()
	}
	this.ServeJSON()
}

func init() {

}
