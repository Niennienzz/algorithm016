package solution

// There are N students in a class.
// Some of them are friends, while some are not.
// Their friendship is transitive in nature.
// For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C.
// And we defined a friend circle is a group of students who are direct or indirect friends.
//
// Given a N*N matrix representing the friend relationship between students in the class.
// If matrix[i][j] = 1, then the ith and jth students are direct friends with each other, otherwise not.
// And you have to output the total number of friend circles among all the students.
//
// Examples:
//
// Input:
// [
//   [1,1,0],
//   [1,1,0],
//   [0,0,1]
// ]
// Output: 2
// Explanation:
//  - The 0th and 1st students are direct friends, so they are in a friend circle.
//  - The 2nd student himself is in a friend circle.
//  - So return 2.
//
// Input:
// [
//   [1,1,0],
//   [1,1,1],
//   [0,1,1]
// ]
// Output: 1
// Explanation:
//  - The 0th and 1st students are direct friends, the 1st and 2nd students are direct friends.
//  - So the 0th and 2nd students are indirect friends.
//  - All of them are in the same friend circle, so return 1.

// 并查集或 DFS 均可. 此处使用并查集, 非常简练.
func findCircleNum(matrix [][]int) int {
	n := len(matrix)
	set := NewDisjointSet(n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if matrix[i][j] == 1 {
				set.Union(i, j)
			}
		}
	}
	return set.Size()
}
