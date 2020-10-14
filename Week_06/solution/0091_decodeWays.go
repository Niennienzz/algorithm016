package solution

import (
	"strconv"
)

// A message containing letters from A-Z is being encoded to numbers using the following mapping:
//
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
//
// Given a non-empty string containing only digits, determine the total number of ways to decode it.
// The answer is guaranteed to fit in a 32-bit integer.
//
// Examples:
//
// Input: s = "12"
// Output: 2
// Explanation: It could be decoded as "AB" (1 2) or "L" (12).
//
// Input: s = "226"
// Output: 3
// Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
//
// Input: s = "0"
// Output: 0
// Explanation:
//  - There is no character that is mapped to a number starting with '0'.
//  - We cannot ignore a zero when we face it while decoding.
//  - So, each '0' should be part of "10" --> 'J' or "20" --> 'T'.
//
// Input: s = "1"
// Output: 1

func numOfDecodingWays(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i < len(dp); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		twoDigits, _ := strconv.Atoi(s[i-2 : i])
		if twoDigits >= 10 && twoDigits <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}
