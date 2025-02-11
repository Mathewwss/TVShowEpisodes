package utils

func SliceMap[T any](list []T, apply func(T) T) []T {
	// 1. Make output
	// 2. Update values
	// 3. Show results

	// 1
	res := make([]T, len(list))

	// 2
	var fn func(int)
	fn = func(i int) {
		if i >= len(list) {
			return
		}

		res[i] = apply(list[i])

		fn(i + 1)
	}
	fn(0)

	// 3
	return res
}
