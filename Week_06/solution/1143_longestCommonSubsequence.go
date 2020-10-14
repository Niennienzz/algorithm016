package solution

// Given two strings text1 and text2, return the length of their longest common subsequence (共同子序列).
// A subsequence of a string is a new string generated from the original string with some characters(can be none)
// deleted without changing the relative order of the remaining characters.
// For instance, "ace" is a subsequence of "abcde" while "aec" is not.
// A common subsequence of two strings is a subsequence that is common to both strings.
// If there is no common subsequence, return 0.
//
// Examples:
//
// Input: text1 = "abcde", text2 = "ace"
// Output: 3
// Explanation: The longest common subsequence is "ace" and its length is 3.
//
// Input: text1 = "abc", text2 = "abc"
// Output: 3
// Explanation: The longest common subsequence is "abc" and its length is 3.
//
// Input: text1 = "abc", text2 = "def"
// Output: 0
// Explanation: There is no such common subsequence, so the result is 0.

// 如何转换为最优子结构?
//  - text1 = "", text2 = 任意; 则为 0.
//  - text1 = "A", text2 = 任意; 则看 "A" 是否存在于 text2 中.
//  - text1 = "...A", text1 = "...A"; 说明最后一个字符是共同子序列, 答案 +1; 并且需要看前面是否也存在共同子序列.
// 经验: 两个字符串的变换问题, 可以转换为一个二维矩阵问题:
//    A B A Z D C
//  B 0 1 1 1 1 1
//  A 1 1 2 2 2 2
//  C 1 1 2 2 2 3 <- LCS("ABAZDC", "BAC")  因为最后一个字符一样, 所以取 3 = 1 + LCS("ABAZD", "BA")
//  B 1 2 2 2 2 3 <- LCS("ABAZDC", "BACB") 因为最后一个字符不同, 所以取 3 = max(LCS("ABAZDC", "BAC"), LCS("ABAZD", "BACB"))
//  A 1 2 3 3 3 3
//  D 1 2 3 3 3 4 <- 答案
func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}

	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
