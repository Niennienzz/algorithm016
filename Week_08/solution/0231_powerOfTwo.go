package solution

// Given an integer n, write a function to determine if it is a power of two.
//
// Examples:
//
// Input: n = 1
// Output: true
//
// Input: n = 16
// Output: true
//
// Input: n = 3
// Output: false
//
// Input: n = 4
// Output: true
//
// Input: n = 5
// Output: false

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	// 清零最低位的 1 之后必为 0.
	// 因为 2 的幂的二进制表示只有一个 1.
	return n&(n-1) == 0
}
