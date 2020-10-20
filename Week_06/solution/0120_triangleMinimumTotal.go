package solution

// Given a triangle, find the minimum path sum from top to bottom.
// Each step you may move to adjacent numbers on the row below.
//
// Examples:
//
// Input:
// [
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
// ]
// Output: The minimum path sum from top to bottom is 11 (i.e. 2 + 3 + 5 + 1 = 11).

// 注意 DP 方程的推导.
// 当前格子的和 = 当前格子的值 + min(下一层对应格子[i+1][j], 下一层斜格子[i+1][j+1]).
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	// 做一个拷贝以免影响原数组, 且新数组可以直接累加.
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
		copy(dp[i], triangle[i])
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] += minInt(dp[i+1][j], dp[i+1][j+1])
		}
	}

	return dp[0][0]
}

// 优化: 可以用一维数组储存 DP 方程.
// 每一层复用并覆盖上一层的值.
// 注意 i 的初始条件, 一维数组的长度等细节.
func minimumTotal2(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i])-1; j++ {
			dp[j] = triangle[i][j] + minInt(dp[j], dp[j+1])
		}
	}

	return dp[0]
}
