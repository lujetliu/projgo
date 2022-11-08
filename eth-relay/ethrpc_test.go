package main

import (
	"fmt"
	"testing"
)

var infuraNodeUrl = "https://mainnet.infura.io/v3/a863a357d92641fcaa7f794b3f81cf7d"
var txHash = "0x53c5b03e392d6aa68a0df26b6d466ae8fbd1c2c5b74f9baae05434ec9a18a282"

func TestNewETHRPCClient(t *testing.T) {
	client := NewETHRPCClient("www.baidu.com").GetRpc()
	if client == nil {
		fmt.Println("初始化失败")
	}

	// client2 := NewETHRPCClient("123://456").GetRpc()
	// if client2 == nil {
	// 	fmt.Println("初始化失败")
	// }
}

func TestGetTransactionByHash(t *testing.T) {
	if txHash == "" || len(txHash) != 66 {
		fmt.Println("非法的交易哈希值")
		return
	}

	txInfo, err := NewETHTPCRequester(infuraNodeUrl).GetTransactionByHash(txHash)
	if err != nil {
		fmt.Printf("查询交易失败: %s\n", err)
		return
	}
	fmt.Println(txInfo)
}
