package orm

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id         int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	LoginName  string    `orm:"size(100)"`
	Password   string    `orm:"size(100)"`
	Name       string    `orm:"size(100)"`
	Mobile     string    `orm:"size(100)"`
	Email      string    `orm:"size(100)"`
	Status     int8      `orm:"NOT NULL;default(1)"`
	IsValid    bool      `orm:"NOT NULL;default(1);type(bit)"`
}

type Role struct {
	Id         int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	RoleName   string    `orm:"size(100)"`
	Remark     string    `orm:"size(100)"`
	Status     int8      `orm:"NOT NULL;default(1)"`
	IsValid    bool      `orm:"NOT NULL;default(1)"`
}

type UserRole struct {
	Id         int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	UserId     int
	RoleId     int
	Status     int8 `orm:"NOT NULL;default(1)"`
	IsValid    bool `orm:"NOT NULL;default(1)"`
}

type Menu struct {
	Id         int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	MenuName   string    `orm:"size(100)"`
	MenuUrl    string    `orm:"size(100)"`
	Level      int       `orm:"default(1)"`
	ParentId   int       `orm:"null"`
	Order      int       `orm: "NOT NULL"`
	Remark     string    `orm:"size(100)"`
	Status     int8      `orm:"NOT NULL;default(1)"`
	IsValid    bool      `orm:"NOT NULL;default(1)"`
}

type RoleMenu struct {
	Id         int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	RoleId     int
	MenuId     int
	Status     int8 `orm:"NOT NULL;default(1)"`
	IsValid    bool `orm:"NOT NULL;default(1)"`
}

type Permission struct {
	Id         int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	PmsName    string    `orm:"size(100)"`
	PmsCode    string    `orm:"size(100)"`
	Remark     string    `orm:"size(100)"`
	Status     int8      `orm:"NOT NULL;default(1)"`
	IsValid    bool      `orm:"NOT NULL;default(1)"`
}

type RolePermission struct {
	Id         int
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
	RoleId     int
	PmsId      int
	Status     int8 `orm:"NOT NULL;default(1)"`
	IsValid    bool `orm:"NOT NULL;default(1)"`
}

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/gomin?charset=utf8&loc=Asia%2FShanghai", 30)

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Role))
	orm.RegisterModel(new(UserRole))
	orm.RegisterModel(new(Menu))
	orm.RegisterModel(new(RoleMenu))
	orm.RegisterModel(new(Permission))
	orm.RegisterModel(new(RolePermission))

	orm.RunSyncdb("default", false, true)
}

//func main() {
//	orm := orm.NewOrm()
//	user := User{Name: "jadson"}
//	err := orm.Read(&user, "name")
//	if err == nil {
//		fmt.Println(user)
//	} else {
//	}
//
//	role := Role{Id: 1}
//	err = orm.Read(&role)
//	fmt.Println(role)
//}
