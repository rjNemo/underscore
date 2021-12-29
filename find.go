package underscore

import "errors"

// Find looks through each value in the slice, returning the first one that passes a truth test (predicate),
// or the default value for the type and an error if no value passes the test.
// The function returns as soon as it finds an acceptable element, and doesn't traverse the entire slice.
func Find[T any](values []T, predicate func(T) bool) (res T, err error) {
	for _, v := range values {
		if predicate(v) {
			return v, nil
		}
	}
	return res, errors.New("value not found")
}
