package utils

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

func Stoi(str string) int {

	i, _ := strconv.Atoi(str)

	return i
}

func Stoi64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func I64tos(i int64) string {
	return fmt.Sprintf("%v", i)
}

func Stod(s string) decimal.Decimal {
	d, _ := decimal.NewFromString(s)
	return d
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
