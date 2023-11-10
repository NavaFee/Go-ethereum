package main

import (
	"fmt"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
)

func EthToTron(str string) string {
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
	a := EthToTron("0x3B85EcD59BCF16246D3400411c8aE43431f9fdb6")
	//TFPwJf9LAkomTJtV4hBh3BjfspUbUJBYjq
	fmt.Println("Tron", a)
	fmt.Println("=====================================")
	b := TronToEth("TFPwJf9LAkomTJtV4hBh3BjfspUbUJBYjq")
	fmt.Println("ETH", b)
}
