package go_wallet

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Println("is valid %v\n", re.MatchString("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")) // true
	fmt.Println("is valid %v\n", re.MatchString("0xXYZ2d8f8d7b001e72cb26a73e4fa1806a51ac79d")) // false

	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Token (ZRX) smart contract address
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)

	}
	isContract = len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract) // is contract: false
}
