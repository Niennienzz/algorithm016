package solution

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
// Examples:
//
// Input: n = 3
// Output:
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	generateParenthesisLoop(0, 0, n, "", &ans)
	return ans
}

// 左括号: 只要数量不超过 n, 随时可以加.
// 右括号: 左个数 > 右个数, 则可以加.
func generateParenthesisLoop(left, right, n int, curr string, ans *[]string) {
	// Recursion terminator.
	if left == n && right == n {
		*ans = append(*ans, curr)
	}

	// Process logic in the current level: append either "(" or ")".
	// Drill down: increase either left or right count, go to the next level.
	if left < n {
		generateParenthesisLoop(left+1, right, n, curr+"(", ans)
	}
	if left > right {
		generateParenthesisLoop(left, right+1, n, curr+")", ans)
	}
}
