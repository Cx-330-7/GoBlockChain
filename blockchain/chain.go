package blockchain

import "fmt"

type Chain struct {
	Blocks []Block
}

// NewChain 创建一个新的区块链
func NewChain() Chain {
	chain := Chain{}
	chain.Blocks = append(chain.Blocks, NewBlock("Genesis Block", ""))
	return chain
}

// GetLastBlock 获取最后一个块
func (c *Chain) GetLastBlock() Block {
	return c.Blocks[len(c.Blocks)-1]
}

// AddBlockToChain 向区块链中加入新块
func (c *Chain) AddBlockToChain(newBlock Block) {
	lastBlock := c.GetLastBlock()
	newBlock.PreviousHash = lastBlock.Hash
	newBlock.Hash = newBlock.ComputeHash()
	c.Blocks = append(c.Blocks, newBlock)
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
