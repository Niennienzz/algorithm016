package solution

// 傻递归, 指数级.
// 每一层的状态数节点数量翻倍, 1x2x2x2...x2 = O(2^N).
func fibNaive(n int) int {
	if n <= 1 {
		return n
	}
	return fibNaive(n-1) + fibNaive(n-2)
}

// 动态规划:
//  - 自顶向下.
//  - 加缓存, 记忆化搜索.
func fibTopDown(n int) int {
	memo := make([]int, n+1)
	return fibTopDownLoop(n, memo)
}

func fibTopDownLoop(n int, memo []int) int {
	if n <= 1 {
		return n
	}
	if memo[n] == 0 {
		memo[n] = fibTopDownLoop(n-1, memo) + fibTopDownLoop(n-2, memo)
	}
	return memo[n]
}

// 动态规划:
//  - 自底向上.
//  - 加缓存, 注意更优化的方式可以只保留前两个值.
func fibBottomUp(n int) int {
	if n <= 1 {
		return n
	}
	memo := make([]int, n+1)
	memo[0], memo[1] = 0, 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
