package service

import (
	"gomin/orm"
)

func MenuList(offset, limit int) (int64, []*orm.Menu) {
	return orm.MenuList(offset, limit)
}

func MenuAdd(menu orm.Menu) int64  {
	return orm.MenuInsert(menu)
}
