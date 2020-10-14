package solution

// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time.
// The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// Now consider if some obstacles are added to the grids.
// How many unique paths would there be?
//
// Examples:
//
// Input:
// [
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
// ]
// Output: 2
// Explanation:
// There is one obstacle in the middle of the 3x3 grid above.
// There are two ways to reach the bottom-right corner:
// 1. Right -> Right -> Down -> Down
// 2. Down -> Down -> Right -> Right

// 动态规划: 最优子结构 -> 一个格子只有两个地方可以到达, 既从其上方或者左侧.

// 自顶向下递归.
func uniquePathsWithObstaclesTopDown(obstacleGrid [][]int) int {
	row := len(obstacleGrid)
	if row == 0 {
		return 0
	}
	col := len(obstacleGrid[0])
	if col == 0 {
		return 0
	}

	countGrid := make([][]*int, row)
	for i := 0; i < row; i++ {
		countGrid[i] = make([]*int, col)
	}

	return uniquePathsWithObstaclesTopDownLoop(row-1, col-1, obstacleGrid, countGrid)
}

func uniquePathsWithObstaclesTopDownLoop(row int, col int, obstacleGrid [][]int, countGrid [][]*int) int {
	if row >= 0 && col >= 0 && countGrid[row][col] != nil {
		return *countGrid[row][col]
	} else if row < 0 || col < 0 || obstacleGrid[row][col] == 1 {
		return 0
	} else if row == 0 && col == 0 {
		return 1
	}

	fromTop := uniquePathsWithObstaclesTopDownLoop(row-1, col, obstacleGrid, countGrid)
	fromLeft := uniquePathsWithObstaclesTopDownLoop(row, col-1, obstacleGrid, countGrid)
	sum := fromTop + fromLeft
	countGrid[row][col] = &sum
	return *countGrid[row][col]
}

// 自底向上.
func uniquePathsWithObstaclesBottomUp(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	result := make([][]int, row)
	for i := range result {
		result[i] = make([]int, col)
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	} else {
		result[0][0] = 1
	}

	for i := 1; i < row; i++ {
		if obstacleGrid[i][0] == 1 || result[i-1][0] == 0 {
			result[i][0] = 0
		} else {
			result[i][0] = 1
		}
	}

	for j := 1; j < col; j++ {
		if obstacleGrid[0][j] == 1 || result[0][j-1] == 0 {
			result[0][j] = 0
		} else {
			result[0][j] = 1
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] != 1 {
				result[i][j] = result[i-1][j] + result[i][j-1]
			} else {
				result[i][j] = 0
			}
		}
	}

	return result[row-1][col-1]
}
