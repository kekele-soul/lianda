package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"lianda/blockchain"
	"lianda/db_mysql"
	_ "lianda/routers"
)

func main() {
	// 实例化一个区块链实例
	bc := blockchain.NewBlockChain()
	fmt.Printf("最新区块的hash值:%x\n",bc.LastHash)

	blocks :=bc.QueryAllBlocks()
	if len(blocks) == 0{
		fmt.Println("暂未查询到区块链数据")
	}
	for _, block := range blocks{
		fmt.Printf("高度：%d\n 哈希：%x\n 上一个哈希：%x\n",block.Height,block.Hash, block.PrevHash)
	}
	fmt.Println()
	return

	//链接数据库
	db_mysql.ConnectDB()
	//
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}


