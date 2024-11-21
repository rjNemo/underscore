package underscore

// Join joins two slices together and returns a Tuple of [T, []P], the selectors allow you to pick the
// keys you want to use from your struct's to join the sets together
func Join[T, P any, S comparable](
	left []T,
	right []P,
	leftSelector func(T) S,
	rightSelector func(P) S) []Tuple[T, []P] {

	var results = make([]Tuple[T, []P], 0, len(left))
	for _, l := range left {
		var matches = Filter(right, func(r P) bool { return leftSelector(l) == rightSelector(r) })
		var tuple = Tuple[T, []P]{Left: l, Right: matches}
		results = append(results, tuple)
	}

	return results
}

// JoinProject joins two slices together and returns a []O where O is defined by the output
// of your projection function
// The selectors allow you to pick the keys from your structure to use as the join keys
// While the projection functions allows you to reformat joined datasets
// (Tuple of [T, []P]) into your own struct or type
func JoinProject[L, R, O any, S comparable](
	left []L,
	right []R,
	leftSelector func(L) S,
	rightSelector func(R) S,
	projection func(Tuple[L, []R]) O) (results []O) {

	for _, x := range Join(left, right, leftSelector, rightSelector) {
		results = append(results, projection(x))
	}

	return results
}
