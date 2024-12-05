package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block 包含的数据有：1.数据 2.之前区块的哈希值 3.自己的哈希值
type Block struct {
	Data         string
	PreviousHash string
	Hash         string
}

// NewBlock 创建一个区块
func NewBlock(data string, previousHash string) Block {
	block := Block{
		Data:         data,
		PreviousHash: previousHash,
	}
	block.Hash = block.ComputeHash()
	return block
}

// ComputeHash 计算当前区块的哈希值
func (b *Block) ComputeHash() string {
	hash := sha256.New()                        // 创建一个新的 SHA-256 哈希计算对象
	hash.Write([]byte(b.Data + b.PreviousHash)) // 将当前区块的数据和前一个区块的哈希值拼接成一个新的字符串，并将其转换为字节切片，然后写入到哈希计算对象中
	return hex.EncodeToString(hash.Sum(nil))
}
