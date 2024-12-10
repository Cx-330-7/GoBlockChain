package main

import (
	"GoBlockChain/blockchain"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	// 初始化区块链
	chain := blockchain.NewChain()

	// 生成公私钥对
	privateKeySender, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKeySender := append(privateKeySender.PublicKey.X.Bytes(), privateKeySender.PublicKey.Y.Bytes()...)

	privateKeyReceiver, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	publicKeyReceiver := append(privateKeyReceiver.PublicKey.X.Bytes(), privateKeyReceiver.PublicKey.Y.Bytes()...)

	// 打印发送者和接收者的公钥
	fmt.Println("Sender Public Key:", hex.EncodeToString(publicKeySender))
	fmt.Println("Receiver Public Key:", hex.EncodeToString(publicKeyReceiver))

	// 创建交易
	tx := blockchain.Transaction{
		From:   hex.EncodeToString(publicKeySender),
		To:     hex.EncodeToString(publicKeyReceiver),
		Amount: 10.5,
	}

	// 签名交易
	err1 := tx.Sign(privateKeySender)
	if err1 != nil {
		return
	}

	// 验证并添加交易
	chain.AddTransaction(tx)
	chain.MineTransactionPool(hex.EncodeToString(publicKeyReceiver))
	chain.Blocks[1].Transactions[0].Signature = "ssdsdw" //篡改交易内容
	fmt.Println(chain.ValidateChain())
}
