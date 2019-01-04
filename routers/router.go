package routers

import (
	"gomin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/test", &controllers.TestController{}, "get:TestApp")
    beego.Router("/table", &controllers.TableController{}, "get:TableList")
    beego.Include(&controllers.CMSController{})

    //  ===============       用户相关    =======================
    beego.Router("/userlisttpl", &controllers.TemplateController{}, "get:UserListTpl")
    beego.Router("/addUserTpl", &controllers.TemplateController{}, "get:AddUserTpl")
    beego.Router("/editUserTpl/?:id", &controllers.TemplateController{}, "get:EditUserTpl")
    beego.Router("/userlist", &controllers.UserController{}, "get:UserList")
    beego.Router("/adduser", &controllers.UserController{}, "post:AddUser")
    beego.Router("/deleteuser", &controllers.UserController{}, "post:DeleteUser")
    beego.Router("/updateUserStatusMulti", &controllers.UserController{}, "post:UpdateUserStatusMulti")
    beego.Router("/queryOneUser/?:id", &controllers.UserController{}, "get:QueryOneUser")
    beego.Router("/updateUser", &controllers.UserController{}, "post:UpdateUser")

    // =================    菜单   =============================
    beego.Router("/menulisttpl", &controllers.TemplateController{}, "get:MenuListTpl")
    beego.Router("/menuaddtpl", &controllers.TemplateController{}, "get:MenuAddTpl")
    beego.Router("/menulist", &controllers.MenuController{}, "get:MenuList")
    beego.Router("/deletemenu", &controllers.MenuController{}, "post:MenuDelete")
    beego.Router("/addmenu", &controllers.MenuController{}, "post:MenuAdd")

}
