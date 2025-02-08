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
