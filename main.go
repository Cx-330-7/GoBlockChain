package main

import (
	"GoBlockChain/blockchain"
	"fmt"
)

func main() {
	// 创建一个区块链
	chain := blockchain.NewChain()

	//创建新区块并添加到区块链
	block1 := blockchain.NewBlock("转账十元", "")
	chain.AddBlockToChain(block1)

	block2 := blockchain.NewBlock("转账三十元", "")
	chain.AddBlockToChain(block2)

	fmt.Println(chain.ValidateChain())
}
