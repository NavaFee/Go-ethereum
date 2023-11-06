package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
)

func main() {

	val := big.NewInt(123)
	com := big.NewInt(456)
	fmt.Println("val", ToDecimal(val, 18))
	fmt.Println("com", ToDecimal(com, 18))
	res := val.Cmp(com)
	fmt.Println(res)

}

func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)
	result = result.Truncate(4)

	return result

}

func BigIntCompare() {

	val := big.NewInt(123)
	com := big.NewInt(456)
	fmt.Println("val", ToDecimal(val, 18))
	fmt.Println("com", ToDecimal(com, 18))
	res := val.Cmp(com)
	fmt.Println(res)

}
