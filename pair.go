package dexsdk

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
)

// Pair trading pair
type Pair struct {
	LiquidityToken Token
	TokenAmounts   [2]TokenAmount
}

// InsufficientAmountError errors
var (
	InsufficientAmountError = errors.New("insufficient amount")
	TokenExistError         = errors.New("pair not this token")
)

// PairAddressCache cache already query pair
var PairAddressCache = map[string]common.Address{}

// NewPair Pair Constructor
func NewPair(token0, token1 TokenAmount) Pair {
	if !token0.Token.SortsBefore(token1.Token) {
		token0, token1 = token1, token0
	}
	return Pair{
		LiquidityToken: NewToken(GetAddress(token0.Token, token1.Token).String(),
			18, token0.ChainId, "xxx-LP", "xxx LPs"),
		TokenAmounts: [2]TokenAmount{token0, token1},
	}
}
func composeKey(token0, token1 Token) string {
	return fmt.Sprintf("%d-%s-%s", token0.ChainId, token0.Address, token1.Address)
}

// GetAddress get Pair address can without Pair Constructor
func GetAddress(tokenA, tokenB Token) common.Address {
	var token0, token1 Token
	if tokenA.SortsBefore(tokenB) {
		token0, token1 = tokenA, tokenB
	} else {
		token0, token1 = tokenB, tokenA
	}
	var key = composeKey(token0, token1)

	if _, ok := PairAddressCache[key]; !ok {
		var salt [32]byte
		hash := crypto.Keccak256(
			token0.ToAddress().Bytes(),
			token1.ToAddress().Bytes(),
		)
		// [32]byte translate
		for i := 0; i < len(salt); i++ {
			salt[i] = hash[i]
		}
		PairAddressCache[key] = crypto.CreateAddress2(common.HexToAddress(FactoryAddressMap[token0.ChainId]), salt, common.FromHex(InitCodeHash[token0.ChainId]))
	}
	return PairAddressCache[key]
}

func (p Pair) Token0() Token {
	return p.TokenAmounts[0].Token
}
func (p Pair) Token1() Token {
	return p.TokenAmounts[1].Token
}
func (p Pair) String() string {
	return p.LiquidityToken.Address
}

func (p Pair) Reverse0() TokenAmount {
	return p.TokenAmounts[0]
}
func (p Pair) Reverse1() TokenAmount {
	return p.TokenAmounts[1]
}
func (p Pair) ReverseOf(token Token) TokenAmount {
	if token.Equals(p.Token0()) {
		return p.Reverse0()
	}
	return p.Reverse1()
}

// GetOutputAmount provide input amount calc output amount
func (p Pair) GetOutputAmount(inputAmount TokenAmount) (TokenAmount, Pair, error) {
	if !p.involvesToken(inputAmount.Token) {
		return TokenAmount{}, Pair{}, TokenExistError
	}
	p.Reverse0().Raw.Cmp(BigZero)
	if p.Reverse0().Raw.Cmp(BigZero) == 0 || p.Reverse1().Raw.Cmp(BigZero) == 0 {
		return TokenAmount{}, Pair{}, InsufficientAmountError
	}
	var outputToken = p.Token1()
	if !inputAmount.Token.Equals(p.Token0()) {
		outputToken = p.Token0()
	}
	var inputReverse = p.ReverseOf(inputAmount.Token)
	var outputReverse = p.ReverseOf(outputToken)

	var inputAmountWithFee = inputAmount.Raw.Mul(FeesNumerator)
	var numerator = inputAmountWithFee.Mul(outputReverse.Raw)
	var denominator = inputAmountWithFee.Add(inputReverse.Raw.Mul(FeesDenominator))
	var outputAmount = NewTokenAmount(outputToken, numerator.Div(denominator))
	return outputAmount, NewPair(inputReverse.Add(inputAmount), outputReverse.Sub(outputAmount)), nil
}

// GetInputAmount provide output amount calc input amount
func (p Pair) GetInputAmount(outputAmount TokenAmount) (TokenAmount, Pair, error) {
	if !p.involvesToken(outputAmount.Token) {
		return TokenAmount{}, Pair{}, TokenExistError
	}
	if p.Reverse0().Raw.Cmp(BigZero) == 0 ||
		p.Reverse1().Raw.Cmp(BigZero) == 0 ||
		outputAmount.Raw.Cmp(p.ReverseOf(outputAmount.Token).Raw) > -1 {
		return TokenAmount{}, Pair{}, InsufficientAmountError
	}
	var inputToken = p.Token1()
	if !outputAmount.Equals(p.Token0()) {
		inputToken = p.Token0()
	}
	var outputReverse = p.ReverseOf(outputAmount.Token)
	var inputReverse = p.ReverseOf(inputToken)
	var numerator = FeesDenominator.Mul(inputReverse.Raw.Mul(outputAmount.Raw))
	var denominator = FeesNumerator.Mul(outputReverse.Raw.Sub(outputAmount.Raw))
	var inputAmount = NewTokenAmount(inputToken, BigOne.Add(numerator.Div(denominator)))
	return inputAmount, NewPair(inputReverse.Add(inputAmount), outputReverse.Sub(outputAmount)), nil
}

// GetLiquidityMinted calc lp has been minted
func (p Pair) GetLiquidityMinted(totalSupply, tokenAmountA, tokenAmountB TokenAmount) (TokenAmount, error) {
	var liquidity decimal.Decimal
	if !tokenAmountA.SortsBefore(tokenAmountB.Token) {
		tokenAmountA, tokenAmountB = tokenAmountB, tokenAmountA
	}
	if totalSupply.Raw.Cmp(BigZero) == 0 {
		sf, _ := p.TokenAmounts[0].Raw.Mul(p.TokenAmounts[1].Raw).Float64()
		liquidity = decimal.NewFromFloat(math.Sqrt(sf)).Sub(MinimumLiquidity)
	} else {

		amount0 := p.TokenAmounts[0].Raw.Div(totalSupply.Raw).Div(p.Reverse0().Raw)
		amount1 := p.TokenAmounts[1].Raw.Div(totalSupply.Raw).Div(p.Reverse1().Raw)
		if amount0.Cmp(amount1) > -1 {
			liquidity = amount0
		} else {
			liquidity = amount1
		}
	}
	if liquidity.Cmp(BigZero) < 1 {
		return TokenAmount{}, InsufficientAmountError
	}
	return NewTokenAmount(p.LiquidityToken, liquidity), nil
}

func (p Pair) involvesToken(token Token) bool {
	return token.Equals(p.Token0()) || token.Equals(p.Token1())
}

// GetLiquidityAmount get liquidity amount
func (p Pair) GetLiquidityAmount(token Token, totalSupply *big.Int, liquidity TokenAmount, feeOn bool, kLast *big.Int) (TokenAmount, error) {
	var TotalSupply = decimal.NewFromBigInt(totalSupply, 0)

	if !p.involvesToken(token) {
		return TokenAmount{}, errors.New("pair no contains this token")
	}
	if !liquidity.Token.Equals(p.LiquidityToken) || !liquidity.Raw.LessThanOrEqual(TotalSupply) {
		return TokenAmount{}, errors.New("liquidity token is wrong")
	}
	var totalSupplyAdjusted decimal.Decimal
	if !feeOn {
		totalSupplyAdjusted = TotalSupply
	} else {
		var kLastDecimal = decimal.NewFromBigInt(kLast, 0)
		if !kLastDecimal.Equal(BigZero) {
			var rootK = p.Reverse0().Raw.Mul(p.Reverse1().Raw)
			var rootKLast = decimal.NewFromBigInt(kLast.Sqrt(kLast), 0)
			if rootK.GreaterThan(rootKLast) {
				var numerator = TotalSupply.Mul(rootK.Sub(rootKLast))
				var denominator = rootKLast.Add(rootK.Mul(BigFive))
				var feeLiquidity = numerator.Div(denominator)
				totalSupplyAdjusted = TotalSupply.Add(feeLiquidity)
			} else {
				totalSupplyAdjusted = TotalSupply
			}

		} else {
			totalSupplyAdjusted = TotalSupply
		}
	}
	return NewTokenAmount(token, liquidity.Raw.Mul(p.ReverseOf(token).Raw).Div(totalSupplyAdjusted)), nil
}
