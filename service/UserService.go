package service

import (
	"gomin/orm"
)

func AddUser(user orm.User) int  {
	id := orm.InsertUser(user)
	return id
}

func DeleteUser(userIds []int)(int64, error)  {
	return orm.DeleteUser(userIds)
}
