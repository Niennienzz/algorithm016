package solution

// Description:
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if key, ok := m[target-v]; ok {
			return []int{key, k}
		} else {
			m[v] = k
		}
	}
	return nil
}
