package orm

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"gomin/model"
)

func MenuList(offset, limit int) (int64, []*Menu) {
	o := orm.NewOrm()
	querySeter := o.QueryTable(new(Menu))
	querySeter = querySeter.Filter("is_valid", 1)
	count, _ := querySeter.Count()
	querySeter = querySeter.Limit(limit, offset)
	menus := make([]*Menu, 0)
	querySeter.All(&menus)
	return count, menus
}

func MenuInsert(menu Menu) int64 {
	menu.Status = 1
	menu.IsValid = true
	count, err := orm.NewOrm().Insert(&menu)
	if err != nil {
		logs.Error("MenuInsert error: ", err)
	}
	return count
}

func MeneDelete(req model.MenuDeleteReq) (int64, error) {
	o := orm.NewOrm()
	count, e := o.QueryTable(new(Menu)).Filter("id__in", req.MenuIds).Update(orm.Params{"is_valid": 0})
	if e != nil {
		logs.Error("MeneDelete error: ", e)
	}
	return count, e
}
