package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	var testnet = "https://bsc-testnet.nodereal.io/v1/c80ff3b41d1c4e25bf779053ca9202a6"

	client, err := ethclient.Dial(testnet)
	if err != nil {
		log.Fatal(err)
	}
	rawTx := "f870f86e81eb8502540be400825208944592d8f8d7b001e72cb26a73e4fa1806a51ac79d880de0b6b3a76400008081e6a01225bcfe364cb0dbba1d479809e2c4419606d4b5551474565d914049dd315d8ba00c4cdc8fd208c9e3c937e218388295b23f737f888876eeea4eb4f756a795ef25"

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
