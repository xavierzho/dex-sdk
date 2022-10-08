package dexsdk

import (
	"github.com/shopspring/decimal"
)

func parseUnits(ether decimal.Decimal, unit int64) decimal.Decimal {
	return ether.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(unit)))
}
func ParseEther(ether string) decimal.Decimal {
	es, _ := decimal.NewFromString(ether)
	return parseUnits(es, 18)
}
func formatUnits(wei decimal.Decimal, unit int64) string {
	return wei.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(unit))).String()
}
func FormatEther(wei decimal.Decimal) string {
	return formatUnits(wei, 18)
}
