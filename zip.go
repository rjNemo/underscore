package underscore

// Zips two slices togther so all the elements of left slice are attached to the corresponding
// elements of the right slice, i.e. [one two three] [1 2 3 4] = [{one, 1} {two, 2} {three, 3}]
// the returned data will be the size of the smallest slice
func Zip[L any, R any](left []L, right []R) []Tuple[L, R] {
	shortest := 0
	if len(left) < len(right) {
		shortest = len(left)
	} else {
		shortest = len(right)
	}

	results := make([]Tuple[L, R], shortest)
	for i := 0; i < shortest; i++ {
		results[i] = Tuple[L, R]{Left: left[i], Right: right[i]}
	}

	return results
}
