package underscore

import (
	"sort"

	"golang.org/x/exp/constraints"
)

// sort any slice ASENDING
func SortSliceASC[T constraints.Ordered](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// sort any slice DESCENDING
func SortSliceDESC[T constraints.Ordered](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}
