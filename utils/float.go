package utils

import (
	"fmt"
)

func Float64ToHex(f float64) (hex string) {
	pref := int64(f / 0x100000000)
	prefStr := fmt.Sprintf("%X", pref)
	suff := int64(f - 0x100000000*float64(pref))
	suff += 0x100000000
	suffStr := fmt.Sprintf("%X", suff)
	suffStr = suffStr[1:len(suffStr)]
	hex = "0x" + prefStr + suffStr
	return
}
