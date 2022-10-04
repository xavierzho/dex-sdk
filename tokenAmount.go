package dexsdk

import "math/big"

type TokenAmount struct {
	Token
	Raw *big.Int
}

func NewTokenAmount(token Token, amount *big.Int) TokenAmount {
	return TokenAmount{
		token,
		amount,
	}
}

func (ta TokenAmount) Add(other TokenAmount) TokenAmount {
	return NewTokenAmount(other.Token, ta.Raw.Add(ta.Raw, other.Raw))
}
func (ta TokenAmount) Sub(other TokenAmount) TokenAmount {
	return NewTokenAmount(other.Token, ta.Raw.Sub(ta.Raw, other.Raw))
}
