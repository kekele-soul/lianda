package main

import (
	"github.com/astaxie/beego"
	"lianda/db_mysql"
	_ "lianda/routers"
)

func main() {
	//链接数据库
	db_mysql.ConnectDB()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()

}

