package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"lianda/blockchain"
	"lianda/db_mysql"
	"lianda/models"
	_ "lianda/routers"
)

func main() {

	user1 := models.User{
		Id:1,
		Phone:"飞飞",
		Password:"1314",
	}
	fmt.Println(user1)
	//json格式
	jsonBytes,_ :=json.Marshal(user1)
	fmt.Println(string(jsonBytes))

	var user2 models.User
	json.Unmarshal(jsonBytes,&user2)
	fmt.Println("反序列化",user2)
	return
	//生成第一个区块
	block := blockchain.NewBlock(0,[]byte{},[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Println(block)
	fmt.Printf("区块的hash值：%x",block.Hash)
	return
	//链接数据库
	db_mysql.ConnectDB()
	//
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}


