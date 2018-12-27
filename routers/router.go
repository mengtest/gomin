package routers

import (
	"gomin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test", &controllers.TestController{}, "get:TestApp")
    beego.Router("/table", &controllers.TableController{}, "get:TableList")
    beego.Include(&controllers.CMSController{})

    beego.Router("/userlisttpl", &controllers.TemplateController{}, "get:UserListTpl")
    beego.Router("/addUserTpl", &controllers.TemplateController{}, "get:AddUserTpl")
    beego.Router("/userlist", &controllers.UserController{}, "get:UserList")
    beego.Router("/adduser", &controllers.UserController{}, "post:AddUser")


}
