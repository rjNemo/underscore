package underscore

// Chunk splits the input slice into groups of size n.
// If n <= 0, it returns nil. The final chunk may be smaller than n.
func Chunk[T any](values []T, n int) [][]T {
	if n <= 0 {
		return nil
	}
	l := len(values)
	if l == 0 {
		return [][]T{}
	}
	chunks := make([][]T, 0, (l+n-1)/n)
	for i := 0; i < l; i += n {
		j := min(i+n, l)
		chunks = append(chunks, values[i:j])
	}
	return chunks
}
