package dexsdk

import (
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// Desc Blockchain Object Description
type Desc struct {
	Name   string `json:"name" yaml:"name"`     // obj name
	Symbol string `json:"symbol" yaml:"symbol"` // obj symbol
}

// Token Represents an ERC20 token with a unique address and some metadata.
type Token struct {
	Desc
	Decimals int8    `json:"decimals" yaml:"decimals"` // token decimals
	Address  string  `json:"address" yaml:"address"`   // token address
	ChainId  ChainId `json:"chainId" yaml:"chainId"`   // token chain id enum
}

// NewToken Token Constructor
func NewToken(address string, decimals int8, chainId ChainId, desc ...string) Token {
	var token Token
	token.Address = address
	token.Decimals = decimals
	token.ChainId = chainId
	if len(desc) >= 2 {
		token.Name = desc[0]
		token.Symbol = desc[1]
	}
	return token
}

// ToAddress convert to ethereum common.Address
func (t Token) ToAddress() common.Address {
	return common.HexToAddress(t.Address)
}

// Equals judgment condition is unique address
func (t Token) Equals(other Token) bool {
	return t.ToAddress() == other.ToAddress()
}

func (t Token) SortsBefore(other Token) bool {
	return strings.ToLower(t.Address) < strings.ToLower(other.Address)
}
