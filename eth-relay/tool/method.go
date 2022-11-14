package tool

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

/*
	要使用"eth_call"来访问智能合约的函数, 必须根据函数的名称生成对应的
	methodId, 可以使用ABI结构体中的Methods成员变量来选出对应的Method对象,
	然后调用Method结构体的ID字段获取methodId
	"methodId"的生成算法比较复杂, 可以直接使用以太坊源码中提供的函数来生成:
	github.com/ethereum/gp-ethereum/accounts/abi/abi.go TODO: 源码
*/

func MakeMethodId(methodName, abiStr string) (string, error) {
	abi := &abi.ABI{} // 实例化"ABI"结构体对象指针
	err := abi.UnmarshalJSON([]byte(abiStr))
	if err != nil {
		return "", err
	}

	// 根据 methodName 获取对应的Method对象
	method := abi.Methods[methodName]
	methodIdBytes := method.ID                         // 获取函数ID
	methodId := "0x" + common.Bytes2Hex(methodIdBytes) // 使用 fmt 输出十六进制
	return methodId, nil
}
