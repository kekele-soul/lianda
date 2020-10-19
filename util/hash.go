package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

func Md5hashSring (data string)(string) {
	hashMd5 :=md5.New()
	hashMd5.Write([]byte(data))
	passwordBytes := hashMd5.Sum(nil)
	return hex.EncodeToString(passwordBytes)
}
func  Md5hashReader(reader io.Reader) (string,error) {
	bytes, err :=ioutil.ReadAll(reader)
	fmt.Println("要计算的hash是：",bytes)
	if err != nil {
		return "",err
	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes),nil
}
/**
 *对数据进行哈希
 */
func  SHA256hashBlock(Data []byte)([]byte)  {
	//对block字段进行拼接
	//对拼接后的数据进行sha256
	sha256hash := sha256.New()
	sha256hash.Write([]byte(""))
	return sha256hash.Sum(nil)

}