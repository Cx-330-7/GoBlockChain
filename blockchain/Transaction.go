package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Transaction struct {
	From      string  // 发送方地址
	To        string  //接收方地址
	Amount    float64 //转账金额
	Signature string  //签名
}

// ComputeHash 计算交易的哈希值
func (t *Transaction) ComputeHash() string {
	data := t.From + t.To + fmt.Sprintf("%.8f", t.Amount)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// Sign 对交易进行签名
func (t *Transaction) Sign(privateKey *ecdsa.PrivateKey) error {
	hash := t.ComputeHash()
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, []byte(hash))
	if err != nil {
		return err
	}
	signature := fmt.Sprintf("%s%s", r.Text(16), s.Text(16))
	t.Signature = signature
	return nil
}
func (t *Transaction) IsValid() bool {
	if t.From == "" {
		//挖矿奖励交易，无需验证
		return true
	}
	// 从公钥获取并创建 ECDSA 公钥
	publicKeyBytes, _ := hex.DecodeString(t.From)
	publicKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     new(big.Int).SetBytes(publicKeyBytes[:len(publicKeyBytes)/2]),
		Y:     new(big.Int).SetBytes(publicKeyBytes[len(publicKeyBytes)/2:]),
	}
	// 验证签名
	hash := t.ComputeHash()
	signature, _ := hex.DecodeString(t.Signature)
	r := new(big.Int).SetBytes(signature[:len(signature)/2])
	s := new(big.Int).SetBytes(signature[len(signature)/2:])
	return ecdsa.Verify(publicKey, []byte(hash), r, s)
}
