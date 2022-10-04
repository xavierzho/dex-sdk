package main

import (
	"fmt"
	dexsdk "github.com/Jonescy/dex-sdk"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
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
	//cli,err := ethclient.Dial("<wss url here!>")

	// new on chain fetcher
	fetcher := dexsdk.NewFetcher(cli)
	token0 := dexsdk.NewToken("0x55d398326f99059fF775485246999027B3197955", 18, dexsdk.BscMain, "Binance-Peg BSC-USD", "BSC-USD")
	token1 := dexsdk.NewToken("0x841E34bc5E80f9cA93aB7b6E3bA85eF4E2A5680C", 18, dexsdk.BscMain, "a cow", "NLGN")
	// get pair reverses of token1 and token0
	pair, err := fetcher.GetReverses(token0, token1)
	if err != nil {
		panic(err)
	}
	// buy some token1

	inputAmount := dexsdk.NewTokenAmount(token1, dexsdk.BigOne.Mul(dexsdk.BigTen, dexsdk.BigOne.Exp(dexsdk.BigTen, big.NewInt(int64(token0.Decimals)), nil)))
	ouptAmount, newPair, err := pair.GetOutputAmount(inputAmount)
	if err != nil {
		panic(err)
	}
	// calc after pair
	fmt.Println(newPair)
	// output amount
	fmt.Println(ouptAmount)
}
