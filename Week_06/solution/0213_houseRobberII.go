package solution

// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed.
// All houses at this place are arranged in a circle. (比起198, 增加了环形的条件.)
// That means the first house is the neighbor of the last one.
// Meanwhile, adjacent houses have a security system connected, and it will automatically contact the police
// if two adjacent houses were broken into on the same night.
//
// Given a list of non-negative integers nums representing the amount of money of each house,
// return the maximum amount of money you can rob tonight without alerting the police.
//
// Examples:
//
// Input: nums = [2,3,2]
// Output: 3
// Explanation:
//  - You cannot rob house 1 (money = 2) and then rob house 3 (money = 2), because they are adjacent houses.
//
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation:
//  - Rob house 1 (money = 1) and then rob house 3 (money = 3).
//  - Total amount you can rob = 1 + 3 = 4.
//
// Input: nums = [0]
// Output: 0

// 和 198 一样, 只是多了首尾相连的环形条件.
func robCircle(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return maxInt(nums[0], nums[1])
	}

	// 不抢第一个.
	return maxInt(rob3(nums[1:]), rob3(nums[:n-1]))
}
