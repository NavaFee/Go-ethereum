package go_listen

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
)

var mainnet = "https://indulgent-radial-arrow.bsc.quiknode.pro/09ae79c0998760c98d010f6435903f3caf0671e4/"
var testnet = "https://bsc-testnet.nodereal.io/v1/c80ff3b41d1c4e25bf779053ca9202a6"

var rawurl = testnet

func Checkvalue() {
	// 区块范围
	startBlock := uint64(34644454)
	endBlock := uint64(34645419)

	// 合约地址
	filterAddress1 := common.HexToAddress("0x0000000000000000000000000000000000001000")
	filterAddress2 := common.HexToAddress("0x0000000000000000000000000000000000001002")
	//contractAddress := common.HexToAddress("0x8e7D337F18CEe1C8839156dE4e6b5E0Ca046ACbf")

	// RPC 连接
	client, err := ethclient.Dial(rawurl)
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

			// 判断交易是否为合约调用
			if tx.To() == nil || *tx.To() == filterAddress1 || *tx.To() == filterAddress2 {
				continue
			}

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

func ListenEvent() {
	// 区块范围
	startBlock := uint64(34644452)
	endBlock := uint64(34645452)

	// 合约地址
	contractAddress := common.HexToAddress("0x8e7D337F18CEe1C8839156dE4e6b5E0Ca046ACbf")

	// RPC 连接
	client, err := ethclient.Dial(rawurl)
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
		hexutil.Encode(log.Data)
		Token := new(big.Int).SetBytes(log.Data)
		//dataBytes, _ := hex.DecodeString(Token[2:]) // 0x000000
		//dataInt := new(big.Int).SetBytes(dataBytes)
		//decString := dataInt.String()
		//fmt.Println("转换结果:", decString)

		txhash := log.TxHash.Hex()

		//fmt.Println(log.Topics[0].Hex()) // 事件签名
		fmt.Printf("Block: %d\n", log.BlockNumber)
		fmt.Printf("TxHash: %s\n", txhash) // 交易哈希
		fmt.Printf("From: %s\n", from.Hex())
		fmt.Printf("To: %s\n", to.Hex())
		fmt.Printf("Value: %s\n", Token)
		// Token value can be extracted from log.Data

		fmt.Println("--------------------------------------------------")
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
