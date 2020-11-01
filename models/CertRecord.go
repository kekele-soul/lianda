package models

import (
	"bytes"
	"encoding/gob"
)

type CertRecord struct {
	CertHash []byte //认证文件的sha256 hash值
	CertId   []byte //认证Id
	CertAuthor string //认证人
	Phone    string //联系方式
	AuthorCard string //身份证号
	FileName  string //认证文件的名称
	FileSize int64  //文件的大小
	CertTime int64 //认证的时间
}
/**
 *认证数据的序列化
 */
func (c CertRecord) SeriaLizeRecord() ([]byte,error ) {
	buff := new(bytes.Buffer)
	err :=gob.NewEncoder(buff).Encode(c)
	return buff.Bytes() ,err
}
func DeSeriaLizeRecord(data []byte) (*CertRecord,error){
	var certRecord *CertRecord
	err :=gob.NewDecoder(bytes.NewReader(data)).Decode(certRecord)
	return certRecord, err
}