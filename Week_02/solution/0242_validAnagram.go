package solution

// Given two strings s and t , write a function to determine if t is an anagram of s.
// You may assume the string contains only lowercase alphabets.
// What if the inputs contain unicode characters? How would you adapt your solution to such case?
//
// Examples:
//
// Input: s = "anagram", t = "nagaram"
// Output: true
//
// Input: s = "rat", t = "car"
// Output: false

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]int)
	for k := range s {
		m[s[k]] += 1
		m[t[k]] -= 1
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
