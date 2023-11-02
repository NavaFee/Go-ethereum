package main

import (
	go_listen "go-eth/go-listen"
)

func main() {
	go_listen.Checkvalue()
	//go_wallet.Wallet()

}

//import (
//	"encoding/hex"
//	"fmt"
//	"math/big"
//)
//
//func main() {
//	hexData := "0000000000000000000000000000000000000000000000000000000000002710"
//
//	dataBytes, err := hex.DecodeString(hexData)
//	if err != nil {
//		// 处理解码错误
//		fmt.Println("解码失败:", err)
//		return
//	}
//
//	dataInt := new(big.Int).SetBytes(dataBytes)
//	decString := dataInt.String()
//
//	fmt.Println("转换结果:", decString)
//}
