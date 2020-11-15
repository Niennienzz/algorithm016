package solution

// Given a string and an integer k, you need to reverse the first k characters for every 2k characters
// counting from the start of the string. If there are less than k characters left, reverse all of them.
// If there are less than 2k but greater than or equal to k characters,
// then reverse the first k characters and left the other as original.
//
// Examples:
//
// Input: s = "abcdefg", k = 2
// Output: "bacdfeg"

func reverseStr(s string, k int) string {
	var (
		bytes = []byte(s)
		end   = len(bytes) - 1
	)

	for i := 0; i <= end; i += 2 * k {
		x := i
		y := i + k - 1
		if y > end {
			y = end
		}
		for x < y {
			bytes[x], bytes[y] = bytes[y], bytes[x]
			x++
			y--
		}
	}

	return string(bytes)
}
