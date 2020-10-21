package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"lianda/blockchain"
	"lianda/db_mysql"
	_ "lianda/routers"
)

func main() {
	//生成第一个区块
	block := blockchain.NewBlock(0,[]byte{},[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Println(block)
	fmt.Printf("区块的hash值：%x",block.Hash)
	return
	//链接数据库
	db_mysql.ConnectDB()
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}


