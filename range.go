package underscore

// Creates a sequence of numbers, i.e. u.Range(0, 3) = [0 1 2 3], while u.Range(3, 0) = [3 2 1 0]
func Range(start int, end int) (result []int) {
	if start < end {
		for i := start; i <= end; i++ {
			result = append(result, i)
		}
	} else {
		for i := start; i >= end; i-- {
			result = append(result, i)
		}
	}

	return result
}
