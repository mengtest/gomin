package orm

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TUser struct {
	Id int
	Name string `orm:"size(100)"`
	Age int
}

func init()  {
	// set default database
	orm.Debug= true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/gomin?charset=utf8&loc=Asia%2FShanghai", 30)

	// register model
	//orm.RegisterModel(new(TUser))

	// create table
	orm.RunSyncdb("default", false, true)
}

//func main() {
//	Insert()
//}

func Insert() {
	o:=orm.NewOrm()
	user:=TUser{Name: "jadson", Age: 20}
	o.Insert(&user)
}

func Read() {
	o := orm.NewOrm()
	user := TUser{Id: 1}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name)
	}
}

func ReadOrUpdate(){
	o := orm.NewOrm()
	user := TUser{Name: "slene"}
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	if created, id, err := o.ReadOrCreate(&user, "Name"); err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
		} else {
			fmt.Println("Get an object. Id:", id)
		}
	}

}

func InsertMulti(){
	o := orm.NewOrm()
	users := []TUser{
		{Name: "slene"},
		{Name: "astaxie"},
		{Name: "unknown"},
	}
	successNums, err := o.InsertMulti(100, users)
	fmt.Println(successNums, err)
}

func Update()  {
	o := orm.NewOrm()
	user := TUser{Id: 1}
	if o.Read(&user) == nil {
		user.Name = "MyName"
		if num, err := o.Update(&user); err == nil {
			fmt.Println(num)
		}
	}

}

func Delete() {
	o := orm.NewOrm()
	if num, err := o.Delete(&TUser{Id: 1}); err == nil {
		fmt.Println(num)
	}
}