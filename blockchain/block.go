package blockchain

import (
	"bytes"
	"lianda/util"
	"time"
)

/*
* 	区块结构的定义
 */
type Block struct {
	Height int64 //区块高度
	TimeStamp int64 //时间戳
	Hash []byte  //区块的hash
	Data []byte  //数据
	PrevHash []byte //上一个区块的哈希
	Version string //版本号
	Nonce   int64   //随机数，用于pow工作量证明算法计算
}
/**
 * 新建一个区块实例，并返回该区块
 */
func NewBlock(height int64,data []byte, prevhash []byte)(Block){
	//构建一个block实例，用于生成区块
	block :=Block{
		Height:height,
		TimeStamp:time.Now().Unix(),
		 Data:data,
		PrevHash: prevhash,
		Version: "0x01",
	}
	//为新生成的block，寻找合适的nonce值
	pow :=NewPow(block)
	nonce := pow.Run()

	//将block的nonce设置为找到的合适的nonce数
	block.Nonce = nonce

	//调用uti.SHA256hash进行计算
	heightBytes,_ := util.IntToBytes(block.Height)//转换成切片类型
	timeBytes, _ := util.IntToBytes(block.TimeStamp)//转换成切片类型
	versiomBytes:= util.StringToBytes(block.Version)//转换成切片类型
	nonceBytes,_ := util.IntToBytes(block.Nonce)
	//byres.Join函数，用于[]byte拼接
	blockBytes :=bytes.Join([][]byte{
		heightBytes,
		timeBytes,
		data,
		prevhash,
		versiomBytes,
		nonceBytes,
	}, []byte{})
	block.Hash = util.SHA256hashBlock(blockBytes)
	return block
}