package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"log"
)

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

	account := crypto.Keccak256(publicKeyBytes[1:])

	fmt.Println("publickeybytes================", hexutil.Encode(account))

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

//package main
//
//import (
//	"fmt"
//	"github.com/shopspring/decimal"
//	"math/big"
//)
//
//func main() {
//
//	val := big.NewInt(123)
//	com := big.NewInt(456)
//	fmt.Println("val", ToDecimal(val, 18))
//	fmt.Println("com", ToDecimal(com, 18))
//	res := val.Cmp(com)
//	fmt.Println(res)
//
//}
//
//func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
//	value := new(big.Int)
//	switch v := ivalue.(type) {
//	case string:
//		value.SetString(v, 10)
//	case *big.Int:
//		value = v
//	}
//
//	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
//	num, _ := decimal.NewFromString(value.String())
//	result := num.Div(mul)
//	result = result.Truncate(4)
//
//	return result
//
//}
//
//func BigIntCompare() {
//
//	val := big.NewInt(123)
//	com := big.NewInt(456)
//	fmt.Println("val", ToDecimal(val, 18))
//	fmt.Println("com", ToDecimal(com, 18))
//	res := val.Cmp(com)
//	fmt.Println(res)
//
//}
