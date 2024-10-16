package underscore

// Convert values to pointers
//
// Instead of:
// v = "value"
// MyPointerVar = &v
//
// Use:
// MyPointerVar = ToPointer("value")
func ToPointer[T any](in T) *T {
	return &in
}
