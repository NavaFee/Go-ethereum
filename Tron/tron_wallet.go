package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"

	"math/big"
)

var base58Alphabets = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func main() {

	privateKey, err := crypto.HexToECDSA("427139B43028A492E2705BCC9C64172392B8DB59F3BA1AEDAE41C88924960091")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publickeybytes================", hexutil.Encode(publicKeyBytes))

	fmt.Println("ethereum account ==== ", crypto.PubkeyToAddress(*publicKeyECDSA).Hex())

	account := crypto.Keccak256(publicKeyBytes[1:])
	last20 := account[len(account)-20:]
	append41 := append([]byte{0x41}, last20...)
	fmt.Println("publickeybytes================", hexutil.Encode(append41))

	fmt.Println("publickeybytes================", hexutil.Encode(account))

	res := base58Encode(append41)
	fmt.Println("publickeybytes================", string(res))
	allBytes := append41 // Your bytes here

	fmt.Println("allBytes", hex.EncodeToString(allBytes))
	str, _ := FromHexAddress(hex.EncodeToString(allBytes))

	rootKey := base58.Encode(allBytes)
	fmt.Println("rootKey", rootKey)
	fmt.Println("rootKey", str)

	hexaddr := ToHexAddress(str)
	fmt.Println("hexaddr", hexaddr)
	hexAddress, _ := hex.DecodeString(hexaddr)

	fmt.Println("hexAddress", hexutil.Encode(hexAddress[1:]))
	//fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//fmt.Println(fromAddress)
	//
	//// Tron uses Keccak-256 (SHA3) hash function, not the standard SHA256.
	//keccak256 := sha3.NewKeccak256()
	//_, err = keccak256.Write(publicKeyBytes[1:]) // remove prefix byte
	//if err != nil {
	//	log.Fatal(err)
	//}
	//publicKeySHA3 := keccak256.Sum(nil)
	//
	//// Take the last 20 bytes of the hashed public key.
	//address := publicKeySHA3[len(publicKeySHA3)-20:]
	//
	//// Prepend the address with 0x41 for Tron's main net.
	//tronAddress := append([]byte{0x41}, address...)
	//
	//// Encode the address using base58.
	//encodedAddress := base58.Encode(tronAddress)

	//fmt.Println("Private Key:", hexutil.Encode(privateKey.D.Bytes()))
	//fmt.Println("Tron Address:", encodedAddress)
}

func ToHexAddress(address string) string {
	return hex.EncodeToString(base58Decode([]byte(address)))
}

func FromHexAddress(hexAddress string) (string, error) {
	fmt.Println("hexAddress", hexAddress)
	addrByte, err := hex.DecodeString(hexAddress)

	if err != nil {
		return "", err
	}

	sha := sha256.New()
	sha.Write(addrByte)
	shaStr := sha.Sum(nil)

	sha2 := sha256.New()
	sha2.Write(shaStr)
	shaStr2 := sha2.Sum(nil)

	addrByte = append(addrByte, shaStr2[:4]...)
	fmt.Println(shaStr2[:4])
	fmt.Println("addrByte", hex.EncodeToString(addrByte))
	return string(base58Encode(addrByte)), nil
}

func base58Encode(input []byte) []byte {
	x := big.NewInt(0).SetBytes(input)
	base := big.NewInt(58)
	zero := big.NewInt(0)
	mod := &big.Int{}
	var result []byte
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabets[mod.Int64()])
	}
	reverseBytes(result)
	return result
}

func base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	for _, b := range input {
		charIndex := bytes.IndexByte(base58Alphabets, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}
	decoded := result.Bytes()
	if input[0] == base58Alphabets[0] {
		decoded = append([]byte{0x00}, decoded...)
	}
	return decoded[:len(decoded)-4]
}

func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
