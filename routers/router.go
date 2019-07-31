package routers

import (
	"api-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/select", &controllers.UserController{}, "Post:Select")
	beego.Router("/list", &controllers.UserController{}, "Post:List")
	beego.Router("/add", &controllers.UserController{}, "Post:Add")
	beego.Router("/update", &controllers.UserController{}, "Post:Update")
	beego.Router("/delete", &controllers.UserController{}, "Post:Delete")
}
