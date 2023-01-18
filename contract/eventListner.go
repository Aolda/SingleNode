package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"unsafe"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func ByteToInt(b []byte) int {
	intFromByte := new(big.Int)
	intFromByte.SetBytes(b)
	return int(intFromByte.Uint64())
}

func main() {
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/ad85dca2c4cf4e5eb8e5cc12ea8f479d")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xbe434BF9cD5c9F7B261D68EbE1bdF588fF6F1f7b")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			if vLog.Topics[0] == common.HexToHash("0x8d397347d14d1ad0861879661d8b750cec64835b214eee49d80b6bf6bb5950de") {
				fmt.Println("Listen 'callAolda'")
			}
			var data []([]byte)
			for i := 0; i < len(vLog.Data); i = i + 32 {
				data = append(data, vLog.Data[i:i+32])
			}

			functionNamePointer := ByteToInt(data[0]) / 32
			argsPointer := ByteToInt(data[1]) / 32
			argsNum := ByteToInt(data[argsPointer])

			functionName := BytesToString(data[functionNamePointer+1])

			var args []string
			for i := argsPointer + 1; i < argsPointer+1+argsNum; i++ {
				ptr := ByteToInt(data[i]) / 32
				arg := BytesToString(data[argsPointer+ptr+2])
				args = append(args, arg)
			}
			fmt.Println(functionName)
			fmt.Println(args)
		}
	}
}
