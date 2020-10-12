package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}
//默认显示的页面、登录页面
func (L *LoginController) Post() {
	L.TplName = "login.html"
}