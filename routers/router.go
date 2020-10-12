package routers

import (
	"lianda/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    beego.Router("/register",&controllers.MainController{})
    beego.Router("/login",&controllers.LoginController{})
}
