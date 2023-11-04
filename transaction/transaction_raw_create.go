package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var testnet = "https://eth-sepolia.g.alchemy.com/v2/VeV1OhHMwDzlaErqViVktjCS0GWfte_-"

	client, err := ethclient.Dial(testnet)
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
	fmt.Println(fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000) // in wei (0.001 eth)
	gasLimit := uint64(21000)          // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x149F8721C909824221Ad7Bd9e0BF19C65050FcDE")
	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     data,
	})

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signedTx", signedTx)
	fmt.Println("----original transaction----")
	fmt.Println("transaction hash:", signedTx.Hash().Hex())
	fmt.Println("transaction value:", signedTx.Value().String())
	fmt.Println("transaction gas limit:", signedTx.Gas())
	fmt.Println("transaction fee cap per gas:", signedTx.GasFeeCap())
	fmt.Println("transaction tip cap per gas:", signedTx.GasTipCap())
	fmt.Println("transaction gas price:", signedTx.GasPrice())
	fmt.Println("transaction nonce:", signedTx.Nonce())
	fmt.Println("transaction data:", hex.EncodeToString(signedTx.Data()))
	fmt.Println("transaction to:", signedTx.To().Hex())

	ts := types.Transactions{signedTx}
	rawTxBytes, _ := rlp.EncodeToBytes(ts[0])
	rawTxHex := hex.EncodeToString(rawTxBytes)

	rawTxBytes1, _ := rlp.EncodeToBytes(ts)
	rawTxHex1 := hex.EncodeToString(rawTxBytes1)

	fmt.Println("rawtxhex", rawTxHex)     // f86...772
	fmt.Println("rawtxhex111", rawTxHex1) // f86...772
	tx1 := new(types.Transaction)
	tx2 := new(types.Transactions)

	rlp.DecodeBytes(rawTxBytes, &tx1)
	fmt.Println("tx1", tx1)
	rlp.DecodeBytes(rawTxBytes1, tx2)
	fmt.Println("tx2", tx2)
	fmt.Println("tx2[0]", (*tx2)[0])

	//err = client.SendTransaction(context.Background(), tx1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	err = client.SendTransaction(context.Background(), (*tx2)[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx1)

	fmt.Println("----decode transaction----")
	fmt.Println("transaction hash:", tx1.Hash().Hex())
	fmt.Println("transaction value:", tx1.Value().String())
	fmt.Println("transaction gas limit:", tx1.Gas())
	fmt.Println("transaction fee cap per gas:", tx1.GasFeeCap())
	fmt.Println("transaction tip cap per gas:", tx1.GasTipCap())
	fmt.Println("transaction gas price:", tx1.GasPrice())
	fmt.Println("transaction nonce:", tx1.Nonce())
	fmt.Println("transaction data:", hex.EncodeToString(tx1.Data()))
	fmt.Println("transaction to:", tx1.To().Hex())

}
