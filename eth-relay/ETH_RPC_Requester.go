package main

import (
	"errors"
	"eth-relay/model"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
)

type ETHRPCRequester struct {
	client *ETHRPClient
}

func NewETHRPCRequester(nodeUrl string) *ETHRPCRequester {
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

// 获取最新区块号
func (r *ETHRPCRequester) GetLatestBlockNumber() (*big.Int, error) {
	methodName := "eth_blockNumber"
	number := ""                                     // 存储结果
	err := r.client.client.Call(&number, methodName) // eth_blockNumber 不需要参数
	if err != nil {
		return nil, fmt.Errorf("获取最新区块号失败:! %s", err.Error())
	}

	tenNumber, _ := new(big.Int).SetString(number[2:], 16)
	return tenNumber, nil
}

// 根据区块号获取区块信息
func (r *ETHRPCRequester) GetBlockInfoByNumber(blockNumer *big.Int) (*model.Block, error) {
	// 八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）
	number := fmt.Sprintf("%#x", blockNumer) // 将 big.Int 转换为16进制字符串
	methodName := "eth_getBlockByNumber"

	block := model.Block{}
	err := r.client.client.Call(&block, methodName, number, true)
	if err != nil {
		return nil, fmt.Errorf("get block info failed! %s", err.Error())
	}

	if block.Number == "" {
		return nil, fmt.Errorf("block info is empty %s", blockNumer.String())
	}

	return &block, nil
}

// 根据区块哈希值获取区块信息
func (r *ETHRPCRequester) GetBlockInfoByHash(blockHash string) (*model.Block, error) {
	methodName := "eth_getBlockByHash"
	block := model.Block{}

	err := r.client.client.Call(&block, methodName, blockHash, true)
	if err != nil {
		return nil, fmt.Errorf("get block info failed! %s", err.Error())
	}

	if block.Number == "" {
		return nil, fmt.Errorf("block info is empty %s", blockHash)
	}

	return &block, nil
}
