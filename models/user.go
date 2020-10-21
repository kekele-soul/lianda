package models

import (
	"fmt"
	"lianda/db_mysql"
	"lianda/util"
)

type User struct {
	Id int `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

//type User struct {
//	Id       string
//	Phone    string
//	Password string
//}
//保存用户信息方法
func (u *User) SaveUser()(int64,error) {
	//密码哈希、脱敏
	u.Password = util.Md5hashSring(u.Password)
	//fmt.Printf("电话号码:%s",u.Phone)
	//fmt.Println("密码:",u.Password)
	row, err :=db_mysql.DB.Exec("insert into user (phone,password) values (?,?)",u.Phone,u.Password)
	if err != nil {
		fmt.Println(err)
		return -1,err
	}
	id, err :=row.RowsAffected()
	if err != nil{
		return -2,err
	}
	return id,nil
}
//查询用户信息
func (u User) QuerUser()(*User,error) {
	//密码哈希、脱敏
	u.Password = util.Md5hashSring(u.Password)
	row := db_mysql.DB.QueryRow("select phone from user where phone = ? and password = ?",u.Phone,u.Password)
	var phone string
 err :=	row.Scan(&phone)
	if err != nil  {
		return nil,err
	}
	return &u,nil
}