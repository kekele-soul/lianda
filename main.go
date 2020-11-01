package main

import (
	"github.com/astaxie/beego"
	"lianda/blockchain"
	"lianda/db_mysql"
	_ "lianda/routers"
)

func main() {
	//qrcode.WriteFile("我爱我老婆,老婆万岁。",qrcode.Medium,256,"./love.png")



	blockchain.NewBlockChain()

	//链接数据库
	db_mysql.ConnectDB()
	//
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}


