package controllers

import (

	"bufio"
	"fmt"
	"lianda/util"
	"lianda/models"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type HomeContorller struct {
	beego.Controller
}
//跳转到新增认证文件的页面upload_file.html
func (h *HomeContorller) Get(){
	phone := h.GetString("phone")
	h.Data["Phone"] = phone
	h.TplName = "home.html"
}
func (h *HomeContorller) Post(){
	//获取客户端上传的文件以及其他from的表单的信息
	//标题
	fmt.Println("----------------------------------------------")
	fileTitlet :=h.Ctx.Request.PostFormValue("upload_title")
	phone := h.Ctx.Request.PostFormValue("phone")
	//文件
	file, header,err :=h.GetFile("upload_file")
	if err != nil {
		fmt.Println(err)
		h.Ctx.WriteString("上传错误，请重试！")
		return
	}
	fmt.Println("标题",fileTitlet)
	fmt.Println("名称",header.Filename)
	fmt.Println("大小",header.Size)
	fmt.Println(file)
	//h.Ctx.WriteString("解析到了，文件是:"+header.Filename)

	//将文件保存到本地的目录当中
	upLoaddir := "./static/img/"+header.Filename
	//打开一个文件夹
	sevefile,err :=os.OpenFile(upLoaddir,os.O_RDWR|os.O_CREATE,777)
	writer := bufio.NewWriter(sevefile)
	//复制到本地
	file_size,err :=io.Copy(writer, file)
	if err != nil {
		h.Ctx.WriteString("保存失败，请重试！")
		return
	}
	defer sevefile.Close()
	fmt.Println("拷贝文件的大小",file_size)
	//计算文件的hash
	//fmt.Println("要计算的文件")
	hash,err :=util.Md5hashReader(file)

	record := models.UploadRecord{}
	record.FileName = header.Filename
	record.FileSize = header.Size
	record.FileTilet = fileTitlet
	record.CertTime = time.Now().Unix()
	record.FileCert = hash
	record.Phone = phone
	str, err := record.SeveRecord()
	fmt.Println(str)
	if err != nil{
		//fmt.Println("111111111111111111",err)
		h.Ctx.WriteString("抱歉，认证错误，请重试！")
		return
	}
	//从数据库当中读取phone用户对应的所有认证记录
	records, err := models.QueryRecordByPhone(phone)
	if err != nil {
		h.Ctx.WriteString("抱歉，获取认证数据失败，请重试")
	}
	fmt.Println(records)
	h.Data["Records"] = records
	h.Data["Phone"] = phone
	h.TplName = "list_record.html"

}
