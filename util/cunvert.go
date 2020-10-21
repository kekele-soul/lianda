package util

import (
	"bytes"
	"encoding/binary"
)

//int 转换成 []byte
func IntToBytes(num int64)([]byte,error){
	buff :=new(bytes.Buffer)//缓冲区
	err :=binary.Write(buff, binary.BigEndian,num)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(),nil
}
//string 转换成 []byte
func StringToBytes(st string) []byte{
	return []byte(st)
}