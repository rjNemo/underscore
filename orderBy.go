package underscore

func OrderBy[T any](list []T, predicate func(T, T) bool) []T {
	swaps := true
	var tmp T

	for swaps {
		swaps = false

		for i := 0; i < len(list)-1; i++ {
			if predicate(list[i], list[i+1]) {
				swaps = true
				tmp = list[i]

				list[i] = list[i+1]
				list[i+1] = tmp
			}
		}
	}

	return list
}
