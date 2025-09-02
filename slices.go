package underscore

import (
	"cmp"
	"sort"
)

// SortSliceASC sorts any slice ASCENDING
func SortSliceASC[T cmp.Ordered](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortSliceDESC sorts any slice DESCENDING
func SortSliceDESC[T cmp.Ordered](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}
