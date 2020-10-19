package controllers

import (
	"github.com/astaxie/beego"
	"lianda/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get(){
	c.Data["Website"] = "www.baidu.com"
	c.TplName = "register.html"
}

func (r * MainController) Post() {
	//解析请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("解析页面错误，请重试")
		return
	}
	//保存用户信息到数据库
	_, err =user.SaveUser()
	//返回前端信息，成功跳登录页面、失败跳出错误页面
	if err != nil {
		r.Ctx.WriteString("抱歉用户注册失败，请重试!")
	}
	//用户注册成功、跳转到登录页面
	r.TplName = "login.html"
}





