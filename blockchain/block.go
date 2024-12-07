package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Block 包含的数据有：1.交易数组 2.之前区块的哈希值 3.自己的哈希值 4.随机数 5.时间戳
type Block struct {
	Transactions []Transaction
	PreviousHash string
	Hash         string
	Timestamp    int64
	Nonce        int
}

// NewBlock 创建一个区块
func NewBlock(transaction []Transaction, previousHash string) Block {
	block := Block{
		Transactions: transaction,
		PreviousHash: previousHash,
		Timestamp:    time.Now().Unix(),
		Nonce:        1,
	}
	block.Hash = block.ComputeHash()
	return block
}

// ComputeHash 计算当前区块的哈希值
func (b *Block) ComputeHash() string {
	// 将交易数组序列化为 JSON 字符串
	transactionsJSON, _ := json.Marshal(b.Transactions)
	hash := sha256.New() // 创建一个新的 SHA-256 哈希计算对象
	hash.Write([]byte(string(transactionsJSON) + b.PreviousHash + strconv.Itoa(b.Nonce)))
	return hex.EncodeToString(hash.Sum(nil))
}

// getAnswer 获取满足工作量证明条件的目标 Hash 前缀（如 '000'），difficulty 是难度值，即前导0的个数
func (b *Block) getAnswer(difficulty int) string {
	answer := ""
	for i := 0; i < difficulty; i++ {
		answer += "0"
	}
	return answer
}

// mine 计算符合工作量证明要求的哈希
func (b *Block) mine(difficulty int) {
	b.Hash = b.ComputeHash()
	for {
		// 如果哈希值前缀不符合要求，增加 nonce 并重新计算
		if b.Hash[:difficulty] != b.getAnswer(difficulty) {
			b.Nonce++
			b.Hash = b.ComputeHash()
		} else {
			// 如果符合要求，结束挖矿
			break
		}
	}
	fmt.Println("挖矿结束", b.Hash)
}
