package service

import (
	"gomin/model"
	"gomin/orm"
)

func AddUser(user orm.User) int {
	id := orm.InsertUser(user)
	return id
}

func DeleteUser(userIds []int) (int64, error) {
	return orm.DeleteUser(userIds)
}

func UpdateUserStatusMulti(userIds []int, targetStatus int8) (int64, error) {
	return orm.UpdateUserStatusMulti(userIds, targetStatus)
}

func QueryOneUser(userId int) orm.User {
	return orm.QueryOneUser(userId)
}

func UpdateUser(user orm.User) (int64, error){
	return orm.UpdateUser(user)
}

func MeneDelete(req model.MenuDeleteReq) (int64, error){
	return orm.MeneDelete(req)
}
