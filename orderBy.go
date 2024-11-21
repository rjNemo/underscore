package underscore

// OrderBy orders a slice by a field value within a struct, the predicate allows you
// to pick the fields you want to orderBy. Use > for ASC or < for DESC
//
//	func (left Person, right Person) bool { return left.Age > right.Age }
func OrderBy[T any](list []T, predicate func(T, T) bool) []T {
	swaps := true
	var tmp T

	//todo: replace with a faster algorithm, this one is pretty simple
	for swaps {
		swaps = false

		for i := 0; i < len(list)-1; i++ {
			if predicate(list[i], list[i+1]) {
				swaps = true
				tmp = list[i]

				list[i] = list[i+1]
				list[i+1] = tmp
			}
		}
	}

	return list
}
