package main

import (
	"context"
	"fmt"
	"time"

	dexsdk "github.com/Jonescy/dex-sdk"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

// bsc chain examples

func main() {
	s := time.Now()
	// new ethereum caller
	cli, err := dexsdk.NewCaller("https://bsc-dataseed1.ninicoin.io/")
	chainID, err := cli.ChainID(context.Background())
	if err != nil {
		panic(err)
	}
	// new on chain fetcher
	fetcher := dexsdk.NewFetcher(cli, dexsdk.ChainId(chainID.Int64()))
	token0, err := fetcher.GetTokenInfo(common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"))
	if err != nil {
		panic(err)
	}
	token1, err := fetcher.GetTokenInfo(common.HexToAddress("0x69b14e8D3CEBfDD8196Bfe530954A0C226E5008E"))
	if err != nil {
		panic(err)
	}
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
	fmt.Printf("%s(%s)\n", dexsdk.ParseEther(outputAmount.Raw), token0.Symbol)
	e := time.Now().Sub(s)
	fmt.Println("spend times", e.Seconds())
}
