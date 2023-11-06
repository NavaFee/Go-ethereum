//package main
//
//import (
//	"bytes"
//	"crypto/sha256"
//	"encoding/hex"
//	"math/big"
//)
//
//var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
//
//func ToHexAddress(address string) string {
//	return hex.EncodeToString(base58Decode([]byte(address)))
//}
//
//func FromHexAddress(hexAddress string) (string, error) {
//	addrByte, err := hex.DecodeString(hexAddress)
//	if err != nil {
//		return "", err
//	}
//
//	sha := sha256.New()
//	sha.Write(addrByte)
//	shaStr := sha.Sum(nil)
//
//	sha2 := sha256.New()
//	sha2.Write(shaStr)
//	shaStr2 := sha2.Sum(nil)
//
//	addrByte = append(addrByte, shaStr2[:4]...)
//	return string(base58Encode(addrByte)), nil
//}
//
//func base58Encode(input []byte) []byte {
//	x := big.NewInt(0).SetBytes(input)
//	base := big.NewInt(58)
//	zero := big.NewInt(0)
//	mod := &big.Int{}
//	var result []byte
//	for x.Cmp(zero) != 0 {
//		x.DivMod(x, base, mod)
//		result = append(result, base58Alphabets[mod.Int64()])
//	}
//	reverseBytes(result)
//	return result
//}
//
//func base58Decode(input []byte) []byte {
//	result := big.NewInt(0)
//	for _, b := range input {
//		charIndex := bytes.IndexByte(base58Alphabets, b)
//		result.Mul(result, big.NewInt(58))
//		result.Add(result, big.NewInt(int64(charIndex)))
//	}
//	decoded := result.Bytes()
//	if input[0] == base58Alphabets[0] {
//		decoded = append([]byte{0x00}, decoded...)
//	}
//	return decoded[:len(decoded)-4]
//}
//
//func reverseBytes(data []byte) {
//	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
//		data[i], data[j] = data[j], data[i]
//	}
//}
//func main() {
//	a := ToHexAddress("TDpBe64DqirkKWj6HWuR1pWgmnhw2wDacE")
//	b, err := FromHexAddress("412A2B9F7641D0750C1E822D0E49EF765C8106524B")
//	if err != nil {
//		panic(err)
//	}
//	println(a)
//	println(b)
//
//}

// 0x2a2b9F7641d0750c1E822D0e49Ef765c8106524b
package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"log"
)

func main() {
	// Replace this with your private key
	privateKeyHex := "427139B43028A492E2705BCC9C64172392B8DB59F3BA1AEDAE41C88924960091"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	// Get the public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	// Generate the Tron address from the public key
	tronAddress := address.PubkeyToAddress(*publicKeyECDSA).Hex()

	fmt.Println("Tron Address:", tronAddress)
}
