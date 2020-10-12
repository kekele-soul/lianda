package models

import (
	"crypto/md5"
	"encoding/hex"
	"kekele/db_mysql"
)

type User struct {
	Id       string
	Phone    string
	Password string
}
//保存用户信息方法
func (u User) SaveUser()(int64,error) {
	hashMd5 :=md5.New()
	hashMd5.Write([]byte(u.Password))
	bytes := hashMd5.Sum(nil)
	u.Password = hex.EncodeToString(bytes)

	row, err :=db_mysql.Db.Exec("insert into user(phone,password)" +"value (?,?)","u.phone,password")
	if err != nil {
		return -1,err
	}
	id, err :=row.RowsAffected()
	if err != nil{
		return -1,err
	}
	return id,nil
}