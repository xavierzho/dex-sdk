package dexsdk

import (
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestGetAddress(t *testing.T) {
	token0 := NewToken("0x55d398326f99059fF775485246999027B3197955", 18, BscMain, "Binance-Peg BSC-USD", "BSC-USD")
	token1 := NewToken("0x841E34bc5E80f9cA93aB7b6E3bA85eF4E2A5680C", 18, BscMain, "a cow", "NLGN")

	if common.HexToAddress("0x6d91542a46Bbc340bE316754345DF8859db34DF2") != GetAddress(token0, token1) {
		t.Fatal("can't calc that address")
	}
}
