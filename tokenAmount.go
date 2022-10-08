package dexsdk

import (
	"github.com/shopspring/decimal"
)

type TokenAmount struct {
	Token
	Raw decimal.Decimal
}

func NewTokenAmount(token Token, amount decimal.Decimal) TokenAmount {
	return TokenAmount{
		token,
		amount,
	}
}

func (ta TokenAmount) Add(other TokenAmount) TokenAmount {
	return NewTokenAmount(other.Token, ta.Raw.Add(other.Raw))
}
func (ta TokenAmount) Sub(other TokenAmount) TokenAmount {
	return NewTokenAmount(other.Token, ta.Raw.Sub(other.Raw))
}
