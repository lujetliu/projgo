package tool

import (
	"fmt"
	"testing"
)

func TestUnlockETHWallet(t *testing.T) {
	address := "0xa61e0e74e2fcd8af5fbf4f8fcab83ace7e23d8b7"
	keyDirs := "../keystores"

	// 第一次解锁出错
	err := UnlockETHWallet(keyDirs, address, "xxxx")
	if err != nil {
		fmt.Println("第一次解锁错误: ", err.Error())
	} else {
		fmt.Println("第一次解锁成功")
	}

	// 第二次密码正确, 解锁成功
	err = UnlockETHWallet(keyDirs, address, "xxxxxxxx")
	if err != nil {
		fmt.Println("第二次解锁错误: ", err.Error())
	} else {
		fmt.Println("第二次解锁成功")

	}
}
