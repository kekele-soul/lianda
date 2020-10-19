package routers

import (
	"lianda/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	//注册页面
    beego.Router("/register",&controllers.MainController{})
    //登录页面
    beego.Router("/login",&controllers.LoginController{})
    //直接登录
    beego.Router("/user_login",&controllers.LoginController{})
	//文件上传接口
	beego.Router("/list_record",&controllers.HomeContorller{})
    //新增存证
    beego.Router("/upload_file.html",&controllers.HomeContorller{})
	//查看数据认证的证书
	//beego.Router("./cert_datail.html",&controllers.)

}
