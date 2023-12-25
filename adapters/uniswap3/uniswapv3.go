package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

/*
Transfer (index_topic_1 address from, index_topic_2 address to, uint256 value)
0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef

Mint (address sender, index_topic_1 address owner, index_topic_2 int24 tickLower, index_topic_3 int24 tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde

IncreaseLiquidity (index_topic_1 uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f

DecreaseLiquidity (index_topic_1 uint256 tokenId, uint128 liquidity, uint256 amount0, uint256 amount1)
0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4

Burn (index_topic_1 address owner, index_topic_2 int24 tickLower, index_topic_3 int24 tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c

PoolCreated (index_topic_1 address token0, index_topic_2 address token1, index_topic_3 uint24 fee, int24 tickSpacing, address pool)
0x783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118f

*/

func main() {

	client, err := ethclient.Dial("https://arb-mainnet.g.alchemy.com/v2/GBqVedptl8QRtYLESyyv4esU845GXhkW")
	if err != nil {
		log.Fatal(err)
	}
	// 获取指定txhash的信息
	// 1. Add liquidity（初始化）：https://arbiscan.io/tx/0x11aa5591a8992fe915435c1e98f3e423634fe49acfadc69aeb6af9dc0212e09b
	// 2. Add liquidity：https://arbiscan.io/tx/0x5d563b62114815f3310ed941858600f1e3655a79a1a39f5c87dfcbf359c125f6
	// 3. Remove liquidity：https://arbiscan.io/tx/0xa2a76b730016abf1775f8a1cd1c60dde98d837db4b3b3b7361570e25cb0a9646
	txHash := common.HexToHash("0xa2a76b730016abf1775f8a1cd1c60dde98d837db4b3b3b7361570e25cb0a9646")
	//tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if isPending {
	//	log.Fatal("pending")
	//}
	//fmt.Println(tx)
	// 获取tx的事件日志
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	for _, log := range receipt.Logs {
		// 判断topic[0]是否为IncreaseLiquidity
		if log.Topics[0].Hex() == "0x3067048beee31b25b2f1681f88dac838c8bba36af25bfb2b7cf7473a5847e35f" {
			// 解析事件日志
			var event struct {
				TokenId   string `json:"tokenId"`
				Liquidity string `json:"liquidity"`
				Amount0   string `json:"amount0"`
				Amount1   string `json:"amount1"`
			}
			// 解析tokenId
			event.TokenId = log.Topics[1].Big().String()

			data := common.Bytes2Hex(log.Data)
			liquidity, _ := new(big.Int).SetString(data[:64], 16)
			event.Liquidity = liquidity.String()
			amount0 := new(big.Int).SetBytes(common.FromHex(data[64:128]))
			event.Amount0 = amount0.String()
			amount1 := new(big.Int).SetBytes(common.FromHex(data[128:]))
			event.Amount1 = amount1.String()
			fmt.Println("Pool_id", log.Address)
			fmt.Println("TokenId:", event.TokenId)
			fmt.Println("Liquidity:", event.Liquidity)
			fmt.Println("Amount0:", event.Amount0)
			fmt.Println("Amount1:", event.Amount1)
			fmt.Println("inflow")

			// 再通过tokenId查询position的具体信息
			// amount0_address
			// amount1_address
		}
		//判断topic[0]是否为DecreaseLiquidity
		if log.Topics[0].Hex() == "0x26f6a048ee9138f2c0ce266f322cb99228e8d619ae2bff30c67f8dcf9d2377b4" {
			// 解析事件日志
			var event struct {
				TokenId   string `json:"tokenId"`
				Liquidity string `json:"liquidity"`
				Amount0   string `json:"amount0"`
				Amount1   string `json:"amount1"`
			}
			// 解析tokenId
			event.TokenId = log.Topics[1].Big().String()

			data := common.Bytes2Hex(log.Data)
			liquidity, _ := new(big.Int).SetString(data[:64], 16)
			event.Liquidity = liquidity.String()
			amount0 := new(big.Int).SetBytes(common.FromHex(data[64:128]))
			event.Amount0 = amount0.String()
			amount1 := new(big.Int).SetBytes(common.FromHex(data[128:]))
			event.Amount1 = amount1.String()
			fmt.Println("Pool_id", log.Address)
			fmt.Println("TokenId:", event.TokenId)
			fmt.Println("Liquidity:", event.Liquidity)
			fmt.Println("Amount0:", event.Amount0)
			fmt.Println("Amount1:", event.Amount1)
			fmt.Println("outflow")

			// 再通过tokenId查询position的具体信息
			// amount0_address
			// amount1_address
		}

	}

}
