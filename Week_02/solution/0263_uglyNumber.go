package solution

// Write a program to check whether a given number is an ugly number.
// Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
//  - 1 is typically treated as an ugly number.
//  - Input is within the 32-bit signed integer range: [−2^31,  2^31 − 1].
//
// Examples:
//
// Input: 6
// Output: true
// Explanation: 6 = 2 × 3
//
// Input: 8
// Output: true
// Explanation: 8 = 2 × 2 × 2
//
// Input: 14
// Output: false
// Explanation: 14 is not ugly since it includes another prime factor 7.

func isUgly(num int) bool {
	for _, factor := range []int{2, 3, 5} {
		for ; num > 0 && num%factor == 0; num /= factor {
		}
	}
	return num == 1
}
