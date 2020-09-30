package solution

// Given a 2D grid map of '1's (land) and '0's (water), count the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.
// You may assume all four edges of the grid are all surrounded by water.
//
// Examples:
//
// Input: grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
// ]
// Output: 1
//
// Input: grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
// ]
// Output: 3

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	var (
		ans    int
		n      = len(grid)
		m      = len(grid[0])
		helper = newNumIslandsHelper(n, m)
	)

	// 遍历二维数组, 找到第一个 1 则是一个岛屿.
	// 从这个 1 开始做 DFS 把所有相连的 1 都变成 0.
	// 再看之后会不会出现新的不相连的 1.
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				helper.loop(&grid, i, j)
				ans += 1
			}
		}
	}

	return ans
}

type numIslandsHelper struct {
	n int
	m int
}

func newNumIslandsHelper(n, m int) *numIslandsHelper {
	return &numIslandsHelper{n, m}
}

func (x *numIslandsHelper) loop(grid *[][]byte, i, j int) {
	// 递归的终止条件: i, j 坐标超出地图范围, 或者当前格子不再是 1.
	if i < 0 || j < 0 || i >= x.n || j >= x.m || (*grid)[i][j] != '1' {
		return
	}

	// 处理当前格子.
	(*grid)[i][j] = '0'

	// 向周围 DFS 辐射.
	x.loop(grid, i+1, j)
	x.loop(grid, i-1, j)
	x.loop(grid, i, j+1)
	x.loop(grid, i, j-1)
}
