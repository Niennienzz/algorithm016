package solution

// Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
// You may return the answer in any order.
//
// Examples:
//
// Input: n = 4, k = 2
// Output:
// [
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
// ]
//
// Input: n = 1, k = 1
// Output: [[1]]

func combine(n int, k int) [][]int {
	ans := make([][]int, 0, 100)
	cur := make([]int, 0, k)
	backtrack(1, k, n, &cur, &ans)
	return ans
}

func backtrack(start, k, n int, cur *[]int, ans *[][]int) {
	if len(*cur) == k {
		cpy := make([]int, k)
		copy(cpy, *cur)
		*ans = append(*ans, cpy)
	}

	for i := start; i <= n; i++ {
		*cur = append(*cur, i)
		backtrack(i+1, k, n, cur, ans)
		*cur = (*cur)[:len(*cur)-1]
	}
}
