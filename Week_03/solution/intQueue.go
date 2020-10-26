package solution

type intQueue struct {
	data []int
}

func newIntQueue() *intQueue {
	return &intQueue{
		data: make([]int, 0),
	}
}

func (x intQueue) len() int {
	return len(x.data)
}

func (x intQueue) isEmpty() bool {
	return x.len() == 0
}

func (x *intQueue) push(v int) {
	x.data = append(x.data, v)
}

func (x *intQueue) pop() int {
	v := x.data[0]
	x.data = x.data[1:]
	return v
}
