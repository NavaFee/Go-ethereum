package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type TransferEvent struct {
	BlockNumber           uint64 `json:"block_number"`
	BlockTimestamp        int64  `json:"block_timestamp"`
	CallerContractAddress string `json:"caller_contract_address"`
	ContractAddress       string `json:"contract_address"`
	EventIndex            int    `json:"event_index"`
	EventName             string `json:"event_name"`
	Result                struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Value string `json:"value"`
	} `json:"result"`
	ResultType struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Value string `json:"value"`
	} `json:"result_type"`
	Event         string `json:"event"`
	TransactionID string `json:"transaction_id"`
}

type Response struct {
	Data    []TransferEvent `json:"data"`
	Success bool            `json:"success"`
	Meta    struct {
		At          int64  `json:"at"`
		Fingerprint string `json:"fingerprint"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
		PageSize int `json:"page_size"`
	} `json:"meta"`
}

func main() {
	//blockNumber := uint64(56323693)
	//blockNumber = blockNumber + 1
	//a := strconv.FormatUint(blockNumber, 10)
	//fmt.Println(a)
	//res, err := GetEvents(a)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_ = res
	//currentblock, err := GetBlockNumber()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("currentblock", currentblock)

	//监听波场信息
	// 1. 获取起始的监听区块，和当前最新的区块
	var startBlock uint64
	var endBlock uint64
	currentblock, err := GetBlockNumber()
	if err != nil {
		log.Fatal(err)
	}
	startBlock = currentblock - 10
	endBlock = currentblock
	for {
		// 2. 获取当前最新的区块号
		endBlock, err = GetBlockNumber()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("startBlock:", startBlock, "   endBlock:", endBlock)
		if startBlock > endBlock {
			time.Sleep(3 * time.Second)
			endBlock = endBlock + 1
		}

		for blockNumber := startBlock; blockNumber <= endBlock; blockNumber++ {
			fmt.Println("--------------------------------------------------", blockNumber)
			res, err := GetEvents(fmt.Sprintf("%d", blockNumber))
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range res.Data {
				// 筛选事件
				_ = v
				//fmt.Println("{From", v.Result.From, ",To:", v.Result.To, ",Value", v.Result.Value, "}")
				// 过滤事件并入库 mysql

			}
			startBlock = blockNumber + 1
		}

	}

}

// 获取指定区块的交易事件
func GetEvents(blockNumber string) (Response, error) {

	url := "https://api.trongrid.io/v1/contracts/TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t/events?event_name=Transfer&only_confirmed=true&limit=200&block_number=" + blockNumber
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	req.Header.Add("TRON-PRO-API-KEY", "53a4078c-4964-4187-b3e8-c067a7fbc236")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.trongrid.io")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	var response Response
	json.Unmarshal(body, &response)

	return response, nil

}

// 获取当前最新的区块号
func GetBlockNumber() (uint64, error) {

	url := "https://api.trongrid.io/v1/blocks/latest/events?only_confirmed=true&limit=1"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	req.Header.Add("TRON-PRO-API-KEY", "53a4078c-4964-4187-b3e8-c067a7fbc236")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Host", "api.trongrid.io")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var response Response
	json.Unmarshal(body, &response)
	return response.Data[0].BlockNumber, nil

}
