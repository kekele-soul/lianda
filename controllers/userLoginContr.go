package controllers

import (
	"github.com/astaxie/beego"
	"lianda/models"
)

type LoginController struct {
	beego.Controller
}

func (L *LoginController) Get() {
	L.TplName = "login.html"
}

//用户登录接口
func (L *LoginController) Post() {
	var user models.User
	//fmt.Println("--------------------------------------------------------------------------------------------------")
	err := L.ParseForm(&user)
	if err != nil {
		L.Ctx.WriteString("抱歉，用户信息解析失败，请重试！")
		return
	}
	//查询数据库用户信息
	u,err :=user.QuerUser()
	//fmt.Println("=================================================================")
	if err != nil {
		L.Ctx.WriteString("抱歉，用户登录失败，请重试！")
		return
	}
	//登陆成功、跳转到核心页面
	L.Data["Phone"] = u.Phone
	L.TplName = "home.html"//{{.phone}}
}