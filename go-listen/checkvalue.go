package go_listen

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

func Checkvalue() {
	// 区块范围
	startBlock := uint64(34673427)
	endBlock := uint64(34673456)

	// 合约地址
	//contractAddress := common.HexToAddress("0x8e7D337F18CEe1C8839156dE4e6b5E0Ca046ACbf")

	// RPC 连接
	client, err := ethclient.Dial("https://bsc-testnet.nodereal.io/v1/14be4adfd27045daa81812e8603cbaea")
	if err != nil {
		log.Fatal(err)
	}

	// 遍历区块
	for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
		fmt.Println("--------------------------------------------------", blockNumber)
		// 获取区块
		block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
		if err != nil {
			log.Fatal(err)
		}

		// 遍历区块中的交易
		for _, tx := range block.Transactions() {
			//判断BNB交易
			if tx.Value().String() == "0" {
				continue
			}

			//// 判断交易是否为合约调用
			//if tx.To() == nil || *tx.To() != contractAddress {
			//	continue
			//}

			// 获取发送者地址
			msg, err := client.TransactionSender(context.Background(), tx, block.Hash(), 0)
			if err != nil {
				log.Fatal(err)
			}

			// 处理合约调用交易
			fmt.Printf("Block: %d\n", blockNumber)
			fmt.Printf("TxHash: %s\n", tx.Hash().Hex())
			fmt.Printf("From: %s\n", msg.Hex())
			fmt.Printf("To: %s\n", tx.To().Hex())
			fmt.Printf("Value: %s\n", tx.Value().String())
			fmt.Println("--------------------------------------------------")
		}
	}
}
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result

}
