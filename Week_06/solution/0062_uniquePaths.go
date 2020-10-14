package solution

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time.
// The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// How many possible unique paths are there?
//
// Examples:
//
// Input: m = 3, n = 7
// Output: 28
//
// Input: m = 7, n = 3
// Output: 28
//
// Input: m = 3, n = 2
// Output: 3
// Explanation:
// From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Down -> Down
// 2. Down -> Down -> Right
// 3. Down -> Right -> Down
//
// Input: m = 3, n = 3
// Output: 6

// 动态规划: 最优子结构 -> 一个格子只有两个地方可以到达, 既从其上方或者左侧.

// 自顶向下递归.
func uniquePathsTopDown(m int, n int) int {
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}

	return uniquePathsTopDownLoop(m-1, n-1, grid)
}

func uniquePathsTopDownLoop(m int, n int, grid [][]int) int {
	if m <= 0 || n <= 0 {
		return 1
	}
	if grid[m][n] == 0 {
		grid[m][n] = uniquePathsTopDownLoop(m-1, n, grid) + uniquePathsTopDownLoop(m, n-1, grid)
	}
	return grid[m][n]
}

// 自底向上.
func uniquePathsBottomUp(m int, n int) int {
	grid := make([][]int, m)

	for i := range grid {
		grid[i] = make([]int, n)
		grid[i][0] = 1
	}

	for j := range grid[0] {
		grid[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[m-1][n-1]
}
