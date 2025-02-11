package utils

func SliceFilter[T any](list []T, filter func(T) bool) []T {
	// 1. Make output
	// 2. Filter original
	// 3. Return filtered

	// 1
	res := make([]T, 0)

	// 2
	var fn func(int)
	fn = func(i int) {
		if i >= len(list) {
			return
		}

		if filter(list[i]) {
			res = append(res, list[i])
		}

		fn(i + 1)
	}
	fn(0)

	// 3
	return res
}
