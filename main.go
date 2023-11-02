package main

//
//import (
//	go_listen "go-eth/go-listen"
//)
//
//func main() {
//	go_listen.Checkvalue()
//	//go_wallet.Wallet()
//
//}

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

	// 区块范围
	//startBlock := uint64(34645400) // 指定的起始块
	//blockCountToScan := 5          // 需要扫描的块数

	//var mainnet = "https://indulgent-radial-arrow.bsc.quiknode.pro/09ae79c0998760c98d010f6435903f3caf0671e4/"
	var testnet = "https://bsc-testnet.nodereal.io/v1/c80ff3b41d1c4e25bf779053ca9202a6"

	var rawurl = testnet

	// RPC 连接
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		log.Fatal(err)
	}

	filterAddress1 := common.HexToAddress("0x0000000000000000000000000000000000001000")
	filterAddress2 := common.HexToAddress("0x0000000000000000000000000000000000001002")
	filterAddress3 := common.HexToAddress("0x0000000000000000000000000000000000001003")
	filteraddresses := map[common.Address]bool{
		filterAddress1: true,
		filterAddress2: true,
	}
	fmt.Println(filteraddresses[filterAddress3])
	fmt.Println(client)
	//var maxBlockNumber uint64
	//db.Table("blocks").Select("MAX(block_number)").Scan(&maxBlockNumber)
	//
	//fmt.Println("Max block number:", maxBlockNumber)
	// 无限循环，直到程序中断
	//for {
	//	// 获取当前最新的区块号
	//	latestBlock, err := client.BlockNumber(context.Background())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// 确定需要扫描的结束块号
	//	endBlock := latestBlock - uint64(blockCountToScan)
	//
	//	// 遍历区块
	//	for blockNumber := startBlock; blockNumber > endBlock; blockNumber-- {
	//
	//		fmt.Println("--------------------------------------------------", blockNumber)
	//		// 获取区块
	//		block, err := client.BlockByNumber(context.Background(), new(big.Int).SetUint64(blockNumber))
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		// 遍历区块中的交易
	//		for _, tx := range block.Transactions() {
	//			// 判断BNB交易
	//
	//			if tx.Value().String() == "0" {
	//				continue
	//			}
	//			// 判断创建合约交易和特殊地址
	//			if tx.To() == nil || *tx.To()  filteraddresses {
	//				continue
	//			}
	//
	//			// 获取发送者地址
	//			_, err := client.TransactionSender(context.Background(), tx, block.Hash(), 0)
	//			if err != nil {
	//				log.Fatal(err)
	//			}
	//
	//		}
	//
	//	}
	//}

}
