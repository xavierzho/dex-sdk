package dexsdk

import (
	"github.com/shopspring/decimal"
)

var FactoryAddressMap = map[ChainId]string{
	BscMain: "0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73",
	BscTest: "0x6725f303b657a9451d8ba641348b6761a6cc7a17",
	EthMain: "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f",
}

var InitCodeHash = map[ChainId]string{
	BscMain: "0x00fb7f630766e6a796048ea87d01acd3068e8ff67d078148a3fa3f4a84f69bd5",
	BscTest: "0xd0d4c4cd0848c93cb4fd1f498d7013ee6bfb25783ea21593d5834f5d250ece66",
	EthMain: "0x96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f",
}

var FeesNumerator = decimal.NewFromInt(9975)
var FeesDenominator = decimal.NewFromInt(10000)
var BigOne = decimal.NewFromInt(1)
var BigZero = decimal.NewFromInt(0)
var BigTen = decimal.NewFromInt(10)
var MinimumLiquidity = decimal.NewFromInt(1000)
