package solution

// Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:
//  - Integers in each row are sorted from left to right.
//  - The first integer of each row is greater than the last integer of the previous row.
//
// Examples:
//
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 3
// Output: true
//
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,50]], target = 13
// Output: false
//
// Input: matrix = [], target = 0
// Output: false

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	lo, hi := 0, m*n-1
	var mid int

	for lo <= hi {
		mid = lo + (hi-lo)/2
		if target == matrix[mid/n][mid%n] {
			return true
		} else if target <= matrix[mid/n][mid%n] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return false
}
