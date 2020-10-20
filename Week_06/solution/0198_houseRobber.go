package solution

// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them
// is that adjacent houses have security system connected and it will automatically contact the police
// if two adjacent houses were broken into on the same night.
//
// Given a list of non-negative integers representing the amount of money of each house,
// determine the maximum amount of money you can rob tonight without alerting the police.
//
// Examples:
//
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation:
//  - Rob house 1 (money = 1) and then rob house 3 (money = 3).
//  - Total amount you can rob = 1 + 3 = 4.
//
// Input: nums = [2,7,9,3,1]
// Output: 12
// Explanation:
//  - Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
//  - Total amount you can rob = 2 + 9 + 1 = 12.

// 多了偷/不偷的状态, DP方程需要增加维度.
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([][2]int, n)

	dp[0][0] = 0       // 不偷第一个房子.
	dp[0][1] = nums[0] // 偷第一个房子.

	for i := 1; i < n; i++ {
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]) // 不偷当前房子, 看前一个偷/不偷的最大值.
		dp[i][1] = dp[i-1][0] + nums[i]           // 偷当前房子, 前一个房子不能偷.
	}

	return maxInt(dp[n-1][0], dp[n-1][1])
}

// 简化 DP 方程的维度.
//  - [i] 表示当前房子不论偷或不偷, 能得到的最大值.
//  - 最终类似斐波那契数列.
func rob2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = maxInt(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = maxInt(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

// 从上面可知, 其实只需要保存最近的两个数, 类似斐波那契数列最优解.
func rob3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	x := nums[0]
	y := maxInt(nums[0], nums[1])

	for i := 2; i < n; i++ {
		x, y = y, maxInt(y, x+nums[i])
	}

	return y
}
