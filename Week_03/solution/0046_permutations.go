package solution

// Given a collection of distinct integers, return all possible permutations.
//
// Examples:
//
// Input: [1,2,3]
// Output:
// [
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
// ]

func permute(nums []int) [][]int {
	var ans [][]int
	permuteLoop(nums, nil, &ans)
	return ans
}

func permuteLoop(nums []int, curr []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, curr...))
		return
	}
	for i := 0; i < len(nums); i++ {
		fir, sec := nums[:i], nums[i+1:]
		newNums := append(append([]int{}, fir...), sec...)
		permuteLoop(newNums, append(curr, nums[i]), ans)
	}
}
