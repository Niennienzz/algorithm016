package solution

type stack struct {
	data []int
}

func newStack() *stack {
	return &stack{
		data: make([]int, 0),
	}
}

func (x *stack) push(v int) {
	x.data = append(x.data, v)
}

func (x *stack) pop() int {
	v := x.data[len(x.data)-1]
	x.data = x.data[:len(x.data)-1]
	return v
}

func (x *stack) top() int {
	return x.data[len(x.data)-1]
}

func (x *stack) isEmpty() bool {
	return len(x.data) == 0
}
