package blockchain

import (
	"lianda/util"
	"time"
)

/*
* 	区块结构的定义
 */
type Block struct {
	Height int //区块高度
	TimeStamp int64 //时间戳
	Hash []byte  //区块的hash
	Data []byte  //数据
	PrevHash []byte //上一个区块的哈希
	Version string //版本号
}
/**
 * 新建一个区块实例，并返回该区块
 */
func NewBlock(height int,data []byte, prevhash []byte)(Block){
	block :=Block{
		Height:height+1,
		TimeStamp:time.Now().Unix(),
		 Data:data,
		PrevHash: prevhash,
		Version: "0x01",
	}
	//block.Hash = util.SHA256hashBlock(Block)
	return block
}