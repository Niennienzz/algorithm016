package solution

// Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.
//
// Examples:
//
// s = "leetcode"
// return 0.
//
// s = "loveleetcode"
// return 2.

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v] += 1
	}
	for k, v := range s {
		if count, ok := m[v]; ok && count == 1 {
			return k
		}
	}
	return -1
}
