package main

import "eth-relay/model"

type ETHRPCRequester struct {
	client *ETHRPClient
}

func NewETHTPCRequester(nodeUrl string) *ETHRPCRequester {
	requester := &ETHRPCRequester{}
	requester.client = NewETHRPCClient(nodeUrl)
	return requester
}

func (r *ETHRPCRequester) GetTransactionByHash(txHash string) (model.Transaction, error) {
	methodName := "eth_getTransactionByHash"
	result := model.Transaction{}
	err := r.client.GetRpc().Call(&result, methodName, txHash)
	return result, err
}
