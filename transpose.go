package underscore

// Transpose flips a matrix over its diagonal, swapping rows and columns.
// Returns an empty slice if the input is empty.
// Assumes all rows have the same length (uses the length of the first row).
//
// Example: Transpose([[1,2,3], [4,5,6]]) → [[1,4], [2,5], [3,6]]
func Transpose[T any](matrix [][]T) [][]T {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]T{}
	}

	rows := len(matrix)
	cols := len(matrix[0])
	result := make([][]T, cols)

	for i := range result {
		result[i] = make([]T, rows)
		for j := range matrix {
			result[i][j] = matrix[j][i]
		}
	}

	return result
}
