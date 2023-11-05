package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-eth/store"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/VeV1OhHMwDzlaErqViVktjCS0GWfte_-")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("e0d11070a7d128f6b2e109107d9561e43b514aeb4de1a2bffe2a0f5b69c8fb8f")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, _ := client.ChainID(context.Background())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xe544948a76193b99aBE0083edD0182B1432d4E8A")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//key := [32]byte{}
	//value := [32]byte{}
	//copy(key[:], []byte("foo"))
	//copy(value[:], []byte("bar"))
	//
	//tx, err := instance.SetItem(auth, key, value)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//txHash := tx.Hash()
	//fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	//
	//tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	//
	//for isPending {
	//	time.Sleep(time.Second * 5)
	//	fmt.Println("pending...")
	//	_, isPending, _ = client.TransactionByHash(context.Background(), txHash)
	//}

	result, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result) // "bar"
}
