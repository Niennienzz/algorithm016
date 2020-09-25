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
	ans := [][]int{}
	combineLoop(1, n, k, nil, &ans)
	return ans
}

func combineLoop(s, n, k int, curr []int, ans *[][]int) {
	if k < 0 {
		return
	}
	if k == 0 {
		*ans = append(*ans, append([]int{}, curr...))
		return
	}
	for i := s; i <= n; i++ {
		combineLoop(i+1, n, k-1, append(curr, i), ans)
	}
}
