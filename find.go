package underscore

import "errors"

func Find[T any](values []T, predicate func(T) bool) (res T, err error) {
	for _, v := range values {
		if predicate(v) {
			return v, nil
		}
	}
	return res, errors.New("value not found")
}
