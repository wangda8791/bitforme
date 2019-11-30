package utils

// "github.com/bn_funds/models"

func Extend(slice []interface{}, element interface{}) []interface{} {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}
