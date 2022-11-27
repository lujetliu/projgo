package tool

import (
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

/*
	解锁钱包
	在实现对交易数据进行签名之前, 要先对当前交易中作为交易发起者的地址进行
	解锁操作; 所谓解锁, 就是获取到地址对应的私钥;
	解锁钱包的操作流程是: 将发起地址的"keystore"文件结合当初设置的密码
	解析出私钥, 将私钥数据放在内存中, 待需要对数据进行签名的时候使用;

	go-ethereum 库中的 keystore.go 中提供用于解锁的函数 Unlock, TODO: 源码

*/

// 全局保存解锁成功的钱包map集合变量
var ETHUnlockMap map[string]accounts.Account

// 全局对应keystore实例
var UnlockKs *keystore.KeyStore

// 解锁以太坊钱包, 传入钱包地址和对应的keystore密码
func UnlockETHWallet(keysDir, address, password string) error {
	if UnlockKs == nil {
		UnlockKs = keystore.NewKeyStore(
			// 服务器端存储 keystore 文件的文件夹
			keysDir,
			keystore.StandardScryptN,
			keystore.StandardScryptP)
		if UnlockKs == nil {
			return errors.New("ks is nil")
		}
	}

	unlock := accounts.Account{Address: common.HexToAddress(address)}
	// ks.Unlock 调用 keystore.go 的解锁函数, 解锁出的私钥将存储在其变量中
	if err := UnlockKs.Unlock(unlock, password); nil != err {
		return errors.New("unlock err: " + err.Error())
	}

	if ETHUnlockMap == nil {
		ETHUnlockMap = map[string]accounts.Account{} // map 的零值, 没有使用make
	}

	ETHUnlockMap[address] = unlock // 解锁成功, 存储
	return nil
}
