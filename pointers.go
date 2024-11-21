package underscore

// ToPointer Convert values to pointers
//
// Instead of:
// v := "value"
// MyPointerVar = &v
//
// Or
// v1 := "value1"
// v2 := 100
//
//	obj := Obj{
//		Field1: &v,
//		Field2: &v2,
//	}
//
// Use:
// MyPointerVar = ToPointer("value")
func ToPointer[T any](in T) *T {
	return &in
}
