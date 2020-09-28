package solution

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minInt3(a, b, c int) int {
	return minInt(a, minInt(b, c))
}
