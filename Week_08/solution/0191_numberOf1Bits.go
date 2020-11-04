package solution

// Write a function that takes an unsigned integer and returns the number of '1' bits it has.
// This value is also known as the Hamming weight.
//
// Examples:
//
// Input: n = 00000000000000000000000000001011
// Output: 3
//
// Input: n = 00000000000000000000000010000000
// Output: 1
//
// Input: n = 11111111111111111111111111111101
// Output: 31

func hammingWeight(n uint32) int {
	var res int
	for n > 0 {
		res += 1
		n = n & (n - 1) // 清零最低位的 1
	}
	return res
}
