package model

// TODO: 熟悉以太坊区块相关字段信息
type Block struct {
	Number           string        `json:"number"`           // 区块号
	Hash             string        `json:"hash"`             // 区块的哈希值
	ParentHash       string        `json:"parentHash"`       // 父区块的哈希值
	Nonce            string        `json:"nonce"`            // 区块的序列号
	Sha3Uncles       string        `json:"sha3Uncles"`       // 当前区块如果打包了叔块, 则此值是叔块的sha3加密值
	LogsBloom        string        `json:"logsBloom"`        // 当前区块的布隆过滤器日志, TODO
	TransactionsRoot string        `json:"transactionsRoot"` // 交易默克尔树的根部hash值
	ReceiptsRoot     string        `json:"stateRoot"`        // 收据默克尔树的根部的哈希值
	Miner            string        `json:"miner"`            // 挖出此区块的旷工的以太坊地址值
	Difficulty       string        `json:"difficulty"`       // 区块的难度值(或许叫难度系数?)
	TotalDifficulty  string        `json:"totalDifficulty"`  // 区块所在的链的总难度
	ExtraData        string        `json:"extraData"`        // 区块的附属收据
	Size             string        `json:"size"`             // 区块总数据量的大小
	GasLimit         string        `json:"gasLimit"`         // 区块的 GasLimit, 和交易的GasLimit不一样
	GasUsed          string        `json:"gasUsed"`          // 当前该区块已经打包了的交易的总燃料费
	Timestamp        string        `json:"timestamp"`        // 区块被确认核实的时间戳, 单位为秒
	Uncles           []string      `json:"uncles"`           // 叔块的哈希数组
	Transactions     []interface{} `json:"transactions"`     // 所有被打包了的交易的数组
}
