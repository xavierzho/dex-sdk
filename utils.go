package dexsdk

import (
	"github.com/shopspring/decimal"
)

func parseUnits(ether decimal.Decimal, unit int64) decimal.Decimal {
	return ether.Div(decimal.NewFromInt(10).Pow(decimal.NewFromInt(unit)))
}
func ParseEther(i interface{}) decimal.Decimal {
	switch ether := i.(type) {
	case string:
		es, _ := decimal.NewFromString(ether)
		return parseUnits(es, 18)
	case decimal.Decimal:
		return parseUnits(ether, 18)
	default:
		return decimal.Decimal{}
	}
}
func formatUnits(wei decimal.Decimal, unit int64) string {
	return wei.Mul(decimal.NewFromInt(10).Pow(decimal.NewFromInt(unit))).String()
}
func FormatEther(wei decimal.Decimal) string {
	return formatUnits(wei, 18)
}
