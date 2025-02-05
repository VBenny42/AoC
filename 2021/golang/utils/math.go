package utils

import "golang.org/x/exp/constraints"

func Abs[T constraints.Signed](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

func MaxSlice[T constraints.Ordered](slice []T) T {
	maxVal := slice[0]
	for _, v := range slice {
		maxVal = max(maxVal, v)
	}
	return maxVal
}
