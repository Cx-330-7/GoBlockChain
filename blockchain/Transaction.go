package blockchain

type Transaction struct {
	From   string  // 发送方地址
	To     string  //接收方地址
	Amount float64 //转账金额
}
