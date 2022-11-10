package main

import (
	"eth-relay/model"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type ERC20BalanceRpcReq struct {
	ContractAddress string // 合约的以太坊地址
	UserAddress     string // 用户的以太坊地址
	ContractDecimal int    // 合约所对应代币单位精确到小数点后的位数
	client          *ETHRPClient
}

func NewEERC20BalanceRpcReq(nodeUrl string) *ERC20BalanceRpcReq {
	requester := &ERC20BalanceRpcReq{}
	requester.client = NewETHRPCClient(nodeUrl)
	return requester
}

// 批量查询: 根据以太坊地址数组, 查询ERC20代币的余额
func (r *ERC20BalanceRpcReq) GetERC20Balances(paramArr []ERC20BalanceRpcReq) ([]string, error) {
	methodName := "eth_call"
	methodId := "0x70a08231" // balanceOf 的 methodId, TODO: 如何获取方法id
	rets := []*string{}

	length := len(paramArr)
	reqs := []rpc.BatchElem{}
	for i := 0; i < length; i++ {
		ret := ""
		arg := &model.CallArg{}
		userAddress := paramArr[i].UserAddress

		// 针对访问 balanceOf 时的必需参数, 查询余额不需要燃料费,
		// 此处不需要设置 gas
		arg.To = common.HexToAddress(paramArr[i].ContractAddress)
		// TODO: data 的参数组合格式, 24个0字符
		arg.Data = methodId + "000000000000000000000000" + userAddress[2:]
		arg.Gas = "0x7a1200"

		// 实例化每个 BatchElem
		req := rpc.BatchElem{
			Method: methodName,
			Args:   []interface{}{arg, "latest"},
			Result: &ret,
		}
		reqs = append(reqs, req)
		rets = append(rets, &ret)
	}

	err := r.client.GetRpc().BatchCall(reqs)
	if err != nil {
		return nil, err
	}

	for _, req := range reqs {
		if req.Error != nil {
			return nil, req.Error
		}
	}

	balances := []string{}
	for _, balance := range rets {
		ten, _ := new(big.Int).SetString((*balance)[2:], 16)
		balances = append(balances, ten.String())
	}

	return balances, nil
}
