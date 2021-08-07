package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  //区块创建时的时间戳
	Data          []byte //区块数据
	PrevBlockHash []byte //前一个区块的哈希值
	Hash          []byte //当前区块的哈希值
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
	}
	block.SetHash()
	return block
}

func NewPanGuBlock() *Block {
	return NewBlock("道生一", []byte{})
}
