package main

import (
	"GoBlockChain/blockchain"
	"fmt"
)

func main() {
	// 创建新的区块链
	chain := blockchain.NewChain()

	// 添加交易
	t1 := blockchain.Transaction{From: "addr1", To: "addr2", Amount: 10}
	t2 := blockchain.Transaction{From: "addr2", To: "addr1", Amount: 5}
	chain.AddTransaction(t1)
	chain.AddTransaction(t2)

	// 矿工地址
	minerAddress := "addr3"
	chain.MineTransactionPool(minerAddress)

	// 打印区块链
	for i, block := range chain.Blocks {
		fmt.Printf("区块 %d: %+v\n", i, block)
	}

	// 验证区块链
	if chain.ValidateChain() {
		fmt.Println("区块链验证成功")
	} else {
		fmt.Println("区块链验证失败")
	}
}
