package blockchain

import (
	"bytes"
	"crypto/sha256"
	"lianda/util"
	"math/big"
)

const DIDDDICULTY = 16
/**
 *工作量证明结构体
 */
type ProoFodWork struct {
	//目标值
	Target *big.Int
	//工作量证明算法对应的哪个区块
	Block Block
}
/**
 *实例化一个pow算法实例
 */
func NewPow(block Block) ProoFodWork{
	targrt := big.NewInt(1)
	targrt.Lsh(targrt, 255-DIDDDICULTY)//左移
	pow := ProoFodWork{
		Target: targrt,
		Block: block,
	}
	return pow
}
/**
 *pow算法：寻找符合条件的nonce值
 */

func (p ProoFodWork) Run() (int64) {
	var nonce int64
	//var bigBlock *big.Int//声明
	bigBlock := new(big.Int)//实例化
	for {
		block := p.Block

		heightBytes,_	:=util.IntToBytes(block.Height)
		timeBytes,_ :=util.IntToBytes(block.TimeStamp)
		versionBytes := util.StringToBytes(block.Version)
		nonceBytes,_ := util.IntToBytes(nonce)

		blockBytes :=bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PrevHash,
			versionBytes,
			nonceBytes,
		},[]byte{})
		sha256hash := sha256.New()
		sha256hash.Write(blockBytes)
		block256hash :=sha256hash.Sum(nil)
		//sha256hash(区块+nonce)对应的大整数
		bigBlock:=bigBlock.SetBytes(block256hash)

		if p.Target.Cmp(bigBlock) == 1{//如果满足条件时，退出循环
			break
		}
		nonce++ //不满足条件就给nonce加一继续循环
	}
	return nonce
}