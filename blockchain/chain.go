package blockchain

import "fmt"

type Chain struct {
	Blocks          []Block
	TransactionPool []Transaction //交易池
	Difficulty      int
	MineReward      int // 挖矿奖励
}

// NewChain 创建一个新的区块链
func NewChain() Chain {
	chain := Chain{
		Blocks:     []Block{},
		MineReward: 50,
		Difficulty: 1,
	}
	// 创建创世区块,创世区块中不包含交易信息
	genesisBlock := NewBlock([]Transaction{}, "")
	chain.Blocks = append(chain.Blocks, genesisBlock)
	return chain
}

// AddTransaction 添加交易到交易池
func (c *Chain) AddTransaction(transaction Transaction) {
	c.TransactionPool = append(c.TransactionPool, transaction)
}

// MineTransactionPool 挖矿并奖励
func (c *Chain) MineTransactionPool(minerRewardAddress string) {
	// 给矿工奖励交易,因为是挖矿得到的奖励，所以没有 From，只有 To
	minerRewardTransaction := Transaction{
		From:   "",
		To:     minerRewardAddress,
		Amount: float64(c.MineReward),
	}
	c.TransactionPool = append(c.TransactionPool, minerRewardTransaction)
	// 创建一个包含交易池中所有交易的区块
	newBlock := NewBlock(c.TransactionPool, c.GetLastBlock().Hash)
	// 挖矿
	newBlock.mine(c.Difficulty)
	//添加区块到区块链，并清空交易池
	c.Blocks = append(c.Blocks, newBlock)
	c.TransactionPool = []Transaction{}
}

// GetLastBlock 获取最后一个块
func (c *Chain) GetLastBlock() Block {
	return c.Blocks[len(c.Blocks)-1]
}

// ValidateChain 验证区块链的完整性
func (c *Chain) ValidateChain() bool {
	if len(c.Blocks) == 1 {
		// 只有创世区块时，直接验证创世区块
		return c.Blocks[0].Hash == c.Blocks[0].ComputeHash()
	}
	for i := 1; i < len(c.Blocks); i++ {
		blockToValidate := c.Blocks[i]
		//验证区块数据是否被篡改
		if blockToValidate.Hash != blockToValidate.ComputeHash() {
			fmt.Println("数据篡改")
			return false
		}
		//验证前一个区块的哈希与当前区块的 previousHash 是否一致
		previousBlock := c.Blocks[i-1]
		if blockToValidate.PreviousHash != previousBlock.Hash {
			fmt.Println("前后区块链接断裂")
			return false
		}
	}
	return true
}

/*
// AddBlockToChain 向区块链中加入新块，因为有了MineTransactionPool方法，所以不再需要这个方法了
func (c *Chain) AddBlockToChain(newBlock Block) {
	lastBlock := c.GetLastBlock()
	newBlock.PreviousHash = lastBlock.Hash
	// 因为这里我们需要通过挖矿来计算区块的哈希值，所以这里这条代码需要注释掉，更新为挖矿代码，挖矿代码中会更新哈希值
	//newBlock.Hash = newBlock.ComputeHash()
	newBlock.mine(c.Difficulty) // 挖矿计算 Hash
	c.Blocks = append(c.Blocks, newBlock)
}
*/
