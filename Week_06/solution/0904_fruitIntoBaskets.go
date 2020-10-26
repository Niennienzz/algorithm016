package solution

// In a row of trees, the i-th tree produces fruit with type tree[i].
// You start at any tree of your choice, then repeatedly perform the following steps:
//  - Add one piece of fruit from this tree to your baskets.  If you cannot, stop.
//  - Move to the next tree to the right of the current tree.  If there is no tree to the right, stop.
//
// Note that you do not have any choice after the initial choice of starting tree:
// you must perform step 1, then step 2, then back to step 1, then step 2, and so on until you stop.
//
// You have two baskets, and each basket can carry any quantity of fruit,
// but you want each basket to only carry one type of fruit each.
//
// What is the total amount of fruit you can collect with this procedure?
//
// Examples:
//
// Input: [1,2,1]
// Output: 3
// Explanation: We can collect [1,2,1].
//
// Input: [0,1,2,2]
// Output: 3
// Explanation: We can collect [1,2,2].
// If we started at the first tree, we would only collect [0, 1].
//
// Input: [1,2,3,2,2]
// Output: 4
// Explanation: We can collect [2,3,2,2].
// If we started at the first tree, we would only collect [1, 2].
//
// Input: [3,3,3,1,2,1,1,2,3,3,4]
// Output: 5
// Explanation: We can collect [1,2,1,1,2].
// If we started at the first tree or the eighth tree, we would only collect 4 fruits.

// 滑动窗口, 快慢指针.
func totalFruit(tree []int) int {
	var ans, slow, fast, newCount int

	// Map is fruitType -> count, which is (tree[idx]) -> count.
	counts := make(map[int]int)

	for fast = 0; fast < len(tree); fast++ {
		// Add the current fruit into counts map.
		// If there are more than 2 types of fruits, clear the slow pointer side.
		counts[tree[fast]] += 1
		for len(counts) >= 3 {
			counts[tree[slow]] -= 1
			if counts[tree[slow]] == 0 {
				delete(counts, tree[slow])
			}
			slow += 1
		}

		// Now there should be only 2 types of fruits in the counts map.
		newCount = 0
		for _, v := range counts {
			newCount += v
		}

		ans = maxInt(ans, newCount)
	}

	return ans
}
