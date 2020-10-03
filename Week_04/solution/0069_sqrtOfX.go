package solution

import (
	"math"
)

// Implement int sqrt(int x).
// Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
// Since the return type is an integer, the decimal digits are truncated and the integer part of the result is returned.
//
// Examples:
//
// Input: 4
// Output: 2
//
// Input: 8
// Output: 2
// Explanation:
//  - The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.

// 二分查找.
// y = x^2 (x>0), 抛物线, 在 y 轴右侧单调递增, 上下界.
func mySqrtUsingBinarySearch(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	var (
		left  = int64(1)
		right = int64(x)
		mid   = int64(0) // 零值或任意值.
		upper = int64(x) // 上界不变.
	)

	for left <= right {
		mid = left + (right-left)/2
		if mid*mid > upper {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return int(right) // 注意思考, 为什么是返回右边.
}

func mySqrtUsingNewtonIteration(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	cur, tar := 1.0, float64(x)

	for {
		pre := cur
		cur = (cur + tar/cur) / 2
		if math.Abs(cur-pre) < 1e-6 {
			return int(cur)
		}
	}
}
