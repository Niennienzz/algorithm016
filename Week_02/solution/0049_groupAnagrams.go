package solution

import (
	"sort"
)

// Given an array of strings, group the anagrams together.
// You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
// typically using all the original letters exactly once.
//
// Examples:
//
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// Input: strs = [""]
// Output: [[""]]
//
// Input: strs = ["a"]
// Output: [["a"]]

// Categorize by sorted strings.
// Two strings are anagrams if and only if their sorted strings are equal.
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		runes := sortableRunes(v)
		sort.Sort(runes)
		k := runes.String()
		m[k] = append(m[k], v)
	}
	ans := make([][]string, 0)
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

type sortableRunes []rune

func (s sortableRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortableRunes) Len() int {
	return len(s)
}

func (s sortableRunes) String() string {
	return string(s)
}
