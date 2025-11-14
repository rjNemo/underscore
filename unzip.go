package underscore

// Unzip splits a slice of tuples into two separate slices.
// The inverse operation of Zip.
//
// Example: Unzip([Tuple{1,"a"}, Tuple{2,"b"}]) → ([1,2], ["a","b"])
func Unzip[L, R any](pairs []Tuple[L, R]) ([]L, []R) {
	if len(pairs) == 0 {
		return []L{}, []R{}
	}

	lefts := make([]L, len(pairs))
	rights := make([]R, len(pairs))

	for i, pair := range pairs {
		lefts[i] = pair.Left
		rights[i] = pair.Right
	}

	return lefts, rights
}
