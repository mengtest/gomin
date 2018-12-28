package orm

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func UserList(offset, limit int) (int64, []*User) {
	o := orm.NewOrm()
	querySelecter := o.QueryTable(new(User))
	querySelecter = querySelecter.OrderBy("-create_time");
	count, e := querySelecter.Count()
	if e != nil {
		fmt.Println("select user error")
	}
	querySelecter = querySelecter.Limit(limit, offset)
	data := make([]*User, 0)
	_, err := querySelecter.All(&data)
	fmt.Println("==> user: {}", len(data))
	if err != nil {
		fmt.Println("select user error")
	}
	logs.Debug("queru user list result: ", data)
	return count, data
}

func InsertUser(user User) int {
	user.IsValid = true
	user.Status = 1
	_, e := orm.NewOrm().Insert(&user)
	if e != nil {
		return -1
	}
	return user.Id
}

func DeleteUser(userIds []int) (int64, error) {
	o := orm.NewOrm()
	count, e := o.QueryTable(new(User)).Filter("id__in", userIds).Update(orm.Params{"is_valid": 0})
	if e != nil {
		logs.Error("delete user error: ", e)
	}
	return count, e
}

func UpdateUserStatusMulti(userIds []int, targetStatus int8) (int64, error) {
	o := orm.NewOrm()
	count, e := o.QueryTable(new(User)).Filter("id__in", userIds).Update(orm.Params{"status": targetStatus})
	if e != nil {
		logs.Error("DisableUser user error: ", e)
	}
	return count, e
}

func QueryOneUser(userId int) User {
	o := orm.NewOrm()
	user := User{Id: userId}
	o.Read(&user)
	return user
}

func UpdateUser(user User) (int64, error) {
	count, err := orm.NewOrm().Update(&user, "login_name", "password", "name", "mobile", "email")
	if err != nil {
		logs.Error("update user error:", err)
	}
	logs.Debug("Update end :", count)
	return count, err
}

