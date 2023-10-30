// package main
//
// import (
//
//	"context"
//	"fmt"
//	"log"
//	"math/big"
//
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/ethclient"
//
// )
//
//	func main() {
//		// 区块范围
//		startBlock := uint64(34644452)
//		endBlock := uint64(34645452)
//
//		// 合约地址
//		contractAddress := common.HexToAddress("0x8e7D337F18CEe1C8839156dE4e6b5E0Ca046ACbf")
//
//		// RPC 连接
//		client, err := ethclient.Dial("https://bsc-testnet.nodereal.io/v1/14be4adfd27045daa81812e8603cbaea")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		// 遍历区块
//		for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
//			fmt.Println("--------------------------------------------------", blockNumber)
//			// 获取区块
//			block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
//			if err != nil {
//				log.Fatal(err)
//			}
//
//			// 遍历区块中的交易
//			for _, tx := range block.Transactions() {
//				// 判断交易是否为合约调用
//				if tx.To() == nil || *tx.To() != contractAddress {
//					continue
//				}
//
//				// 获取发送者地址
//				msg, err := client.TransactionSender(context.Background(), tx, block.Hash(), 0)
//				if err != nil {
//					log.Fatal(err)
//				}
//
//				// 处理合约调用交易
//				fmt.Printf("Block: %d\n", blockNumber)
//				fmt.Printf("TxHash: %s\n", tx.Hash().Hex())
//				fmt.Printf("From: %s\n", msg.Hex())
//				fmt.Printf("To: %s\n", tx.To().Hex())
//				fmt.Printf("Value: %s\n", tx.Value().String())
//				fmt.Println("--------------------------------------------------")
//			}
//		}
//	}
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 区块范围
	startBlock := uint64(34644452)
	endBlock := uint64(34645452)

	// 合约地址
	contractAddress := common.HexToAddress("0x8e7D337F18CEe1C8839156dE4e6b5E0Ca046ACbf")

	// RPC 连接
	client, err := ethclient.Dial("https://bsc-testnet.nodereal.io/v1/14be4adfd27045daa81812e8603cbaea")
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(startBlock),
		ToBlock:   new(big.Int).SetUint64(endBlock),
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{
			{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, log := range logs {
		if len(log.Topics) < 3 {
			continue
		}
		from := common.BytesToAddress(log.Topics[1].Bytes())
		to := common.BytesToAddress(log.Topics[2].Bytes())
		fmt.Printf("Block: %d\n", log.BlockNumber)
		fmt.Printf("From: %s\n", from.Hex())
		fmt.Printf("To: %s\n", to.Hex())
		// Token value can be extracted from log.Data
		fmt.Println("--------------------------------------------------")
	}
}
