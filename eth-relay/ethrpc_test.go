package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

var infuraNodeUrl = "https://mainnet.infura.io/v3/a863a357d92641fcaa7f794b3f81cf7d"
var txHash = "0x53c5b03e392d6aa68a0df26b6d466ae8fbd1c2c5b74f9baae05434ec9a18a282"
var wallet = "0xeBec795c9c8bBD61FFc14A6662944748F299cAcf"

// 测试新建 go-ethereum rpc 客户端连接
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

// 测试查询交易信息
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

// 测试批量查询交易信息
func TestGetTransactions(t *testing.T) {
	var txHashs = []string{
		"0x714c0e40cb90e53593c4e1ba0d24078a5033253d5c7eada69f3454ffc1665c2e",
		"0x714c0e40cb90e53593c4e1ba0d24078a5033253d5c7eada69f3454ffc1665c2e",
		"0x366a2932d605e5007aa31428a8ef5a0ee928a4c5b714b87d66c1c776712518f9",
	}
	txInfos, err := NewETHTPCRequester(infuraNodeUrl).GetTransactions(txHashs)
	if err != nil {
		fmt.Printf("批量查询交易失败: %s\n", err)
		return
	}
	txs, _ := json.Marshal(txInfos)
	fmt.Println(string(txs))
}

// 测试查询ETH余额
func TestGetETHBalance(t *testing.T) {
	balance, err := NewETHTPCRequester(infuraNodeUrl).GetETHBalance(wallet)
	if err != nil {
		fmt.Printf("查询ETH余额失败: %s\n", err)
		return
	}
	fmt.Println(balance)
}

// 测试批量查询ETH余额
func TestGetETHBalanceList(t *testing.T) {
	var address = []string{
		wallet,
		"0xdAC17F958D2ee523a2206206994597C13D831ec7",
	}
	balances, err := NewETHTPCRequester(infuraNodeUrl).GetETHBalanceList(address)
	if err != nil {
		fmt.Printf("批量查询ETH余额失败: %s\n", err)
		return
	}
	fmt.Println(balances)
}
