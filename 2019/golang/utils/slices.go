package utils

func RemoveIndices[T any](s []T, indices ...int) (result []T) {
	m := make(map[int]struct{}, len(indices))
	for _, i := range indices {
		m[i] = struct{}{}
	}

	for i, v := range s {
		if _, ok := m[i]; !ok {
			result = append(result, v)
		}
	}

	return
}

func MapSlices[T any, U any](s []T, f func(T) U) (result []U) {
	result = make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
