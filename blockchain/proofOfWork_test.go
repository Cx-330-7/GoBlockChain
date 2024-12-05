package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"
)

// sha256Hash 用于计算并返回输入字符串的SHA256哈希值
func sha256Hash(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// proofOfWork 实现简单的工作量证明算法
func proofOfWork() (string, int) {
	data := "test"
	x := 0
	for {
		// 计算工作量证明，strconv.Itoa(x)是把 x 转换为字符串
		hashValue := sha256Hash(data + strconv.Itoa(x))
		// 判断哈希值的前4个字符是否为'0000'
		if hashValue[:4] != "0000" {
			x++
		} else {
			// 找到符合条件的哈希值
			return hashValue, x
		}
	}
}

// TestProofOfWork 测试工作量证明的函数
func TestProofOfWork(t *testing.T) {
	hash, x := proofOfWork()
	t.Log("Hash found:", hash)
	t.Log("Value of x:", x)
}
