package solution

// Implement pow(x, n), which calculates x raised to the power n (i.e. x^n).
// Constraints:
//  -100.0 < x < 100.0
//  -2^31 <= n <= 2^31 - 1
//  -10^4 <= x^n <= 10^4
//
// Examples:
//
// Input: x = 2.00000, n = 10
// Output: 1024.00000
//
// Input: x = 2.10000, n = 3
// Output: 9.26100
//
// Input: x = 2.00000, n = -2
// Output: 0.25000
// Explanation: (2)^-2 = (1/2)^2 = 1/4 = 0.25

// 偶数指数: 2^10 = 2^5 乘以自己.
// 奇数指数: 2^5 = 2^2 乘以自己, 再乘以底数2.
// 负数指数: 底数变倒数, 指数负号取消.
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	return myPowLoop(x, n)
}

func myPowLoop(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	half := myPowLoop(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
