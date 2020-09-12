package solution

// You are climbing a stair case. It takes n steps to reach to the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Examples:
//
// Input: 2
// Output: 2
// Explanation: There are two ways to climb to the top.
//  - 1 step + 1 step
//  - 2 steps
//
// Input: 3
// Output: 3
// Explanation: There are three ways to climb to the top.
//  - 1 step + 1 step + 1 step
//  - 1 step + 2 steps
//  - 2 steps + 1 step

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	prev1, prev2 := 0, 1
	var sum int
	for i := 2; i <= n; i++ {
		sum = prev1 + prev2
		prev1 = prev2
		prev2 = sum
	}
	return prev2 + prev1
}
