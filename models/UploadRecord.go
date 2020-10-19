package models

import (
	"fmt"
	"lianda/db_mysql"
	"lianda/util"
)

/**
上传文件记录 结构体定义
*/
type UploadRecord struct {
	Id int
	FileName string //名字
	FileSize int64 //大小
	FileCert string//认证号
	FileTilet string //标签名
	CertTime int64 //时间戳
	FormatCertime string//该字段仅在前端展示
	Phone string //对应的用户
}
/**
保存上传记录到数据库
*/
func (u UploadRecord) SeveRecord()(int64, error){
	//fmt.Println("============================")
	fmt.Println(u.FileName)
	fmt.Println(u.FileTilet)
	rs, err :=db_mysql.DB.Exec("insert into upload_record(file_name,file_size,file_cert,file_tilet,file_time,phone)" +
		"value (?,?,?,?,?,?)", u.FileName, u.FileSize, u.FileCert, u.FileTilet, u.CertTime, u.Phone)
	if err != nil {
		return -1,err
	}
	id, err :=rs.RowsAffected()
	if err != nil {
		return -2,err
	}
	return id,nil
}
/*
读取数据库当中phone用户对应的所有认证数据
 */
func QueryRecordByPhone(phone string) ([]UploadRecord, error) {
	//fmt.Println("---------========================")
	res, err :=db_mysql.DB.Query("select id ,file_name,file_size,file_cert,file_tilet,file_time,phone from upload_record where phone = ?",phone)
	if err != nil{
		return nil, err
	}
	//读取
	records := make([]UploadRecord,0)
	for res.Next()  {
		var record UploadRecord
		err :=res.Scan(&record.Id,&record.FileName,&record.FileSize,&record.FileCert,&record.FileTilet,&record.CertTime,&record.Phone)
		if err != nil {
			return nil,err
		}
		//时间转换 record.Certrime
		record.FormatCertime =util.TimeFormat(record.CertTime,0,util.TIME_FORMAT_THREE)
		records = append(records,record)
	}
	return records,nil
}