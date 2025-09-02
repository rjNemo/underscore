package underscore

import "testing"

// Ensure the unexported marker methods are executed for coverage.
func TestResultIsResultMarker(t *testing.T) {
	var ok Ok[int]
	ok.isResult()

	var er Err[int]
	er.isResult()
}
