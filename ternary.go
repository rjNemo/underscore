package underscore

// Ternary returns the first argument if the condition is true, otherwise the second argument.
// Ternary is a special form of the if statement. It allows you to write code that is more concise and less verbose.
func Ternary[T any](condition bool, pos, neg T) T {
	if condition {
		return pos
	}
	return neg
}
