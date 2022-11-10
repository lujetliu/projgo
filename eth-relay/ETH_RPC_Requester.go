package main

import (
	"errors"
	"eth-relay/model"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
)

type ETHRPCRequester struct {
	client *ETHRPClient
}

func NewETHTPCRequester(nodeUrl string) *ETHRPCRequester {
	requester := &ETHRPCRequester{}
	requester.client = NewETHRPCClient(nodeUrl)
	return requester
}

// 查询交易信息
func (r *ETHRPCRequester) GetTransactionByHash(txHash string) (model.Transaction, error) {
	methodName := "eth_getTransactionByHash"
	result := model.Transaction{}
	err := r.client.GetRpc().Call(&result, methodName, txHash)
	return result, err
}

// 批量查询交易信息
func (r *ETHRPCRequester) GetTransactions(txHashs []string) ([]*model.Transaction, error) {
	methodName := "eth_getTransactionByHash"
	rets := []*model.Transaction{}
	size := len(txHashs)

	reqs := []rpc.BatchElem{}
	for i := 0; i < size; i++ {
		ret := model.Transaction{}
		// 实例化每个BatchElem
		req := rpc.BatchElem{
			Method: methodName,
			Args:   []interface{}{txHashs[i]},
			Result: &ret,
		}
		reqs = append(reqs, req)
		rets = append(rets, &ret)
	}
	err := r.client.GetRpc().BatchCall(reqs) // 传入 BatchElem 数组, 发起批量请求
	return rets, err
}

// 查询ETH余额
func (r *ETHRPCRequester) GetETHBalance(address string) (string, error) {
	methodName := "eth_getBalance"
	var balance string
	err := r.client.GetRpc().Call(&balance, methodName, address, "latest")
	if err != nil {
		return "", err
	}
	if balance == "" {
		return "", errors.New("eth balance is null")
	}
	// 查询返回的余额结果是一个十进制的字符串, 转换为十进制数, 并防止数位溢出
	// TODO: 熟悉 big 包的用法及常用场景
	ten, _ := new(big.Int).SetString(balance[2:], 16)
	return ten.String(), err
}

// 批量查询ETH余额
func (r *ETHRPCRequester) GetETHBalanceList(address []string) ([]string, error) {
	methodName := "eth_getBalance"
	rets := []*string{}
	size := len(address)

	reqs := []rpc.BatchElem{}
	for i := 0; i < size; i++ {
		var ret string
		// 实例化每个BatchElem
		req := rpc.BatchElem{
			Method: methodName,
			Args:   []interface{}{address[i], "latest"},
			Result: &ret,
		}
		reqs = append(reqs, req)
		rets = append(rets, &ret)
	}
	err := r.client.GetRpc().BatchCall(reqs) // 传入 BatchElem 数组, 发起批量请求
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
