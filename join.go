package underscore

// Joins two slices together and returns a Tuple of [T, []P], the selectors allow you to pick the
// keys from your structure you would like to join the sets together with
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

// Joins two slices together and returns a []R where R is defined by the output of your projection
// function
// The selectors allow you to pick the keys from your structure you would like to join the sets
// together with.
// While the projection functions allows you to reformat joined datasets contains in the
// Tuple of [T, []P] into your own structure or type
func JoinProject[T, P, R any, S comparable](
	left []T,
	right []P,
	leftSelector func(T) S,
	rightSelector func(P) S,
	projection func(Tuple[T, []P]) R) (results []R) {

	for _, x := range Join(left, right, leftSelector, rightSelector) {
		results = append(results, projection(x))
	}

	return results
}
