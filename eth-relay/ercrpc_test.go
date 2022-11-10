package main

import (
	"fmt"
	"testing"
)

// 测试批量查询ERC20代币余额
func TestGetERC20Balances(t *testing.T) {
	wallet := "0x63076b9603460dfe3Dd7cB2FF43666C2CD83Aa42"
	contract1 := "0xb5CCd87e613379f815E9Ac5De19A92531E8e9050"
	contract2 := "0x2a3b273695a045ec263970e3c86c23800a0f04fc"

	// 构造参数
	params := []ERC20BalanceRpcReq{}
	item := ERC20BalanceRpcReq{
		ContractAddress: contract1,
		UserAddress:     wallet,
		ContractDecimal: 18,
	}
	params = append(params, item)
	item.ContractAddress = contract2
	params = append(params, item)

	balances, err := NewEERC20BalanceRpcReq(infuraNodeUrl).GetERC20Balances(params)
	if err != nil {
		fmt.Printf("批量查询ERC20代币余额失败: %s\n", err)
		return
	}
	fmt.Println(balances)
}
