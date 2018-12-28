package service

import (
	"gomin/orm"
)

func MenuList(offset, limit int) (int64, []*orm.Menu) {
	return orm.MenuList(offset, limit)
}
