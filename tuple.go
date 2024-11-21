package underscore

// Tuple is a generic tuple type.
// It is used to return multiple values from a function.
type Tuple[L, R any] struct {
	Left  L
	Right R
}
