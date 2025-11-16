package underscore

// FoldRight is like Reduce but processes elements from right to left.
// Also known as foldr in Haskell.
//
// Example: FoldRight([]int{1,2,3}, 0, func(n, acc int) int { return n - acc })
//
//	→ 1 - (2 - (3 - 0)) = 1 - (2 - 3) = 1 - (-1) = 2
func FoldRight[T, P any](values []T, acc P, fn func(T, P) P) P {
	for i := len(values) - 1; i >= 0; i-- {
		acc = fn(values[i], acc)
	}
	return acc
}
