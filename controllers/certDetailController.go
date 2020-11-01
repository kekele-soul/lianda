package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"lianda/blockchain"
	"strings"
)

/*
 * 証書詳情信息查看頁面信息
 */
type CertDetaliControoller struct {
	beego.Controller
}
func (c *CertDetaliControoller) Get() {
	//0、获取前端页面get请求是携带的cert_id数据
	CertId := c.GetString("cert_id")
	fmt.Printf("要查询认证号")
	//1、准备数据：根据cert_id到区块链上查询具体的信息
	block,err :=blockchain.CHAIN.QueryBlockByCerId([]byte(CertId))
	if err != nil{
		c.Ctx.WriteString("查询链上数据失败")
		return
	}
	//查询未遇到错误，两种情况，查到和未查到
	if block == nil{
		c.Ctx.WriteString("抱歉，未查询到数据，请重试！")
		return
	}
	//CertId = hex.EncodeToString(block.Data)
	c.Data["CertId"] =strings.ToUpper(string((block.Data)))
	//2、跳转页面
	c.TplName = "cert_datail.html"//跳轉到詳情頁面
}
