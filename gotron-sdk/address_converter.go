package main

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

func EthToTron1(str string) string {
	str = "41" + str[2:]
	a := address.HexToAddress(str)
	return a.String()
}

func TronToEth(str string) string {
	a, _ := address.Base58ToAddress(str)
	b := []byte(a.Hex())
	return "0x" + string(b[4:])
}

func main() {
	a := EthToTron1("0x8a0ad710d07ffec9850800af9e50506b94a83e7b")
	//TFPwJf9LAkomTJtV4hBh3BjfspUbUJBYjq
	fmt.Println("Tron", a)
	fmt.Println("=====================================")
	b := TronToEth("TFPwJf9LAkomTJtV4hBh3BjfspUbUJBYjq")
	fmt.Println("ETH", b)
}
