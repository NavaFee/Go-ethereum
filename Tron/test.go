package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"
)

func ToTronAddress(ethAddress string) (string, error) {
	// 去除以太坊地址前缀 "0x"
	ethAddress = strings.TrimPrefix(ethAddress, "0x")

	// 解码以太坊地址为字节数组
	ethAddressBytes, err := hex.DecodeString(ethAddress)
	if err != nil {
		return "", err
	}

	// 对以太坊地址进行 Keccak-256 哈希
	hash := sha3.NewLegacyKeccak256()
	hash.Write(ethAddressBytes)
	hashedBytes := hash.Sum(nil)

	// 取 Keccak-256 哈希值的后 20 个字节
	tronAddressBytes := hashedBytes[len(hashedBytes)-20:]

	// 添加 Tron 地址前缀
	tronAddressBytesWithPrefix := append([]byte{0x41}, tronAddressBytes...)

	// 对 Tron 地址进行 Base58 编码
	tronAddress := base58.Encode(tronAddressBytesWithPrefix)

	return tronAddress, nil
}

func main() {
	ethAddress := "0x2a2b9f7641d0750c1e822d0e49ef765c8106524b"
	tronAddress, err := ToTronAddress(ethAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Tron Address:", tronAddress)
}
