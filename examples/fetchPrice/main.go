package main

import (
	"fmt"
	dexsdk "github.com/Jonescy/dex-sdk"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"
)

// bsc chain examples

func main() {
	// if not websocket endpoint using rpc url
	rpcCli, err := rpc.Dial("https://bsc-dataseed1.ninicoin.io/")
	if err != nil {
		panic(err)
	}
	cli := ethclient.NewClient(rpcCli)
	// also has websocket endpoint using wss url
	//cli, err = ethclient.Dial("<wss here>")

	// new on chain fetcher
	fetcher := dexsdk.NewFetcher(cli)
	token0 := dexsdk.NewToken("0x55d398326f99059fF775485246999027B3197955", 18, dexsdk.BscMain, "Binance-Peg BSC-USD", "BSC-USD")
	token1 := dexsdk.NewToken("0x69b14e8D3CEBfDD8196Bfe530954A0C226E5008E", 9, dexsdk.BscMain, "a cow", "NLGN")
	// get pair reverses of token1 and token0
	pair, err := fetcher.GetReverses(token0, token1)
	if err != nil {
		panic(err)
	}
	// buy some token1

	inputAmount := dexsdk.NewTokenAmount(token1, decimal.NewFromInt(1000000000).Mul(dexsdk.BigTen.Pow(decimal.NewFromInt(int64(token1.Decimals)))))
	//fmt.Println(inputAmount.Raw)
	outputAmount, newPair, err := pair.GetOutputAmount(inputAmount)
	if err != nil {
		panic(err)
	}
	// calc after pair
	fmt.Println("new pair:", newPair)
	// output amount
	fmt.Println(dexsdk.ParseEther(outputAmount.Raw.String()))
}
