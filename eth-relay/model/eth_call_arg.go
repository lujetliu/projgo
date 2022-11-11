package model

import "github.com/ethereum/go-ethereum/common"

// TODO: 熟悉eth_call入参
type CallArg struct {
	// common.Address 是以太坊依赖包的地址类型, 其原型是[20]byte数组
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Gas       string         `json:"gas"`
	GasPricre string         `json:"gas_pricre"`
	Value     string         `json:"value"`
	Data      string         `json:"data"`
	Nonce     string         `json:"nonce"`
}
