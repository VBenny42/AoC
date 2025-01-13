package utils

import "golang.org/x/exp/constraints"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MaxSlice[T constraints.Ordered](slice []T) T {
	maxVal := slice[0]
	for _, v := range slice {
		maxVal = max(maxVal, v)
	}
	return maxVal
}
