package solution

// Given two words ('beginWord' and 'endWord'), and a dictionary's word list,
// find the length of shortest transformation sequence from beginWord to endWord, such that:
//  - Only one letter can be changed at a time.
//  - Each transformed word must exist in the word list.
//
// Note:
//  - Return 0 if there is no such transformation sequence.
//  - All words have the same length.
//  - All words contain only lowercase alphabetic characters.
//  - You may assume no duplicates in the word list.
//  - You may assume 'beginWord' and 'endWord' are non-empty and are not the same.
//
// Examples:
//
// Input:
//  - beginWord = "hit",
//  - endWord = "cog",
//  - wordList = ["hot","dot","dog","lot","log","cog"]
// Output: 5
//  - As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog".
//
//
// Input:
//  - beginWord = "hit"
//  - endWord = "cog"
//  - wordList = ["hot","dot","dog","lot","log"]
// Output: 0
//  - The endWord "cog" is not in wordList, therefore no possible transformation.

func ladderLength(beginWord string, endWord string, wordList []string) int {
	return 0
}
