package dexsdk

import (
	"github.com/Jonescy/dex-sdk/abi/pair-bsc"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shopspring/decimal"
)

type Fetcher struct {
	backend bind.ContractBackend
}

func NewFetcher(b bind.ContractBackend) Fetcher {
	return Fetcher{
		b,
	}
}

// GetReverses currently support bsc
func (f Fetcher) GetReverses(token0, token1 Token) (Pair, error) {
	if !token0.SortsBefore(token1) {
		token0, token1 = token1, token0
	}
	contract, err := pair.NewPair(GetAddress(token0, token1), f.backend)
	if err != nil {
		return Pair{}, err
	}
	result, err := contract.GetReserves(&bind.CallOpts{})
	if err != nil {
		return Pair{}, err
	}
	reverse0 := decimal.NewFromBigInt(result.Reserve0, 0)
	reverse1 := decimal.NewFromBigInt(result.Reserve1, 0)
	if BigOne.Equal(reverse0) || BigOne.Equal(reverse1) {
		return Pair{}, InsufficientAmountError
	}
	return NewPair(NewTokenAmount(token0, reverse0), NewTokenAmount(token1, reverse1)), nil
}
