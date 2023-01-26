package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"reflect"
	"unsafe"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	compiler "aolda/compiler"
	"github.com/joho/godotenv"
)

func findZeroFromByte(b []byte) (bool, int) {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return true, i
		}
	}
	return false, 0
}

func BytesToString(b []byte) string {
	isExist, index := findZeroFromByte(b)
	var bb []byte
	if isExist {
		bb = b[:index]
	} else {
		bb = b
	}
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bb))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func ByteToInt(b []byte) int {
	intFromByte := new(big.Int)
	intFromByte.SetBytes(b)
	return int(intFromByte.Uint64())
}

func main() {
	err := godotenv.Load(".env")
	apiKey := os.Getenv("INFURA_API_KEY")
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/" + apiKey)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xb3D008f2b892b9476cDEb75F701ee8Fb1E23AFc0")
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
			if vLog.Topics[0] == common.HexToHash("0x8d397347d14d1ad0861879661d8b750cec64835b214eee49d80b6bf6bb5950de") {
				fmt.Println("Listen 'callAolda'")
			}
			var data []([]byte)
			for i := 0; i < len(vLog.Data); i = i + 32 {
				data = append(data, vLog.Data[i:i+32])
			}

			fileNamePointer := ByteToInt(data[0]) / 32
			functionNamePointer := ByteToInt(data[1]) / 32
			argsPointer := ByteToInt(data[2]) / 32
			argsNum := ByteToInt(data[argsPointer])

			fileName := BytesToString(data[fileNamePointer+1])
			functionName := BytesToString(data[functionNamePointer+1])

			var args []string
			for i := argsPointer + 1; i < argsPointer+1+argsNum; i++ {
				ptr := ByteToInt(data[i]) / 32
				arg := BytesToString(data[argsPointer+ptr+2])
				args = append(args, arg)
			}
			res := compiler.ExecuteJS("script.js", fileName, functionName, args)
			fmt.Println(res)
			//SetValue(functionName, args, res)
		}
	}
}
