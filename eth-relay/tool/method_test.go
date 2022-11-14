package tool

import (
	"fmt"
	"testing"
)

func TestMakeMethodId(t *testing.T) {
	// 加法智能合约的ABI数据, TODO: 智能合约ABI
	contractABI := `
		[ { "constant": true, "inputs": [ { "name": "arg1", "type": "uint8" }, { "name":
		"arg2", "type": "uint8" } ], "name": "add", "outputs": [ { "name": "", "type":
		"uint8" } ], "payable": false, "stateMutability": "pure", "type": "function" } ]
	`

	methodName := "add"
	methodId, err := MakeMethodId(methodName, contractABI)
	if err != nil {
		fmt.Println("生成 methodId 失败", err.Error())
		return
	}
	fmt.Println("生成  methodId 成功", methodId)
}
