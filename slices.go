package underscore

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// SortSliceASC sorts any slice ASCENDING
func SortSliceASC[T constraints.Ordered](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortSliceDESC sorts any slice DESCENDING
func SortSliceDESC[T constraints.Ordered](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}
