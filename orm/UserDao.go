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
	_, e := orm.NewOrm().Insert(&user)
	if e != nil {
		return -1
	}
	return user.Id
}

func UpdateUser(user User) int64 {
	count, error := orm.NewOrm().Update(&user)
	if error != nil {
		logs.Error("update user error:", error)
	}
	logs.Debug("Update end :", count)
	return count
}

func init() {
	logs.SetLevel(logs.LevelDebug)
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	logs.EnableFuncCallDepth(true)
}
