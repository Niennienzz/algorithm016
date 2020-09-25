package solution

// Given a set of distinct integers, return all possible subsets (the power set).
// The solution set must not contain duplicate subsets.
//
// Examples:
//
// Input: nums = [1,2,3]
// Output:
// [
//   [3],
//   [1],
//   [2],
//   [1,2,3],
//   [1,3],
//   [2,3],
//   [1,2],
//   []
// ]

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	subsetsLoop(nums, 0, nil, &ans)
	return ans
}

func subsetsLoop(nums []int, index int, curr []int, ans *[][]int) {
	if index == len(nums) {
		*ans = append(*ans, append([]int{}, curr...))
		return
	}

	// Not pick value at current list.
	subsetsLoop(nums, index+1, curr, ans)

	// Pick value at current list.
	curr = append(curr, nums[index])
	subsetsLoop(nums, index+1, curr, ans)
}
