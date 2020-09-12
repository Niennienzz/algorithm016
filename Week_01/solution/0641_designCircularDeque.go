package solution

// Design your implementation of the circular double-ended queue (deque).

type MyCircularDeque struct {
	data []int
	cap  int
	head int
	tail int
}

func NewMyCircularDeque(k int) MyCircularDeque {
	return MyCircularDeque{
		data: make([]int, k+1),
		cap:  k + 1,
		head: 0,
		tail: 0,
	}
}

func (x *MyCircularDeque) InsertFront(value int) bool {
	if x.IsFull() {
		return false
	}
	x.data[x.head] = value
	x.head++
	x.head %= x.cap
	return true
}

func (x *MyCircularDeque) InsertLast(value int) bool {
	if x.IsFull() {
		return false
	}
	x.tail--
	x.tail = (x.tail + x.cap) % x.cap
	x.data[x.tail] = value
	return true
}

func (x *MyCircularDeque) DeleteFront() bool {
	if x.IsEmpty() {
		return false
	}
	x.head--
	x.head = (x.head + x.cap) % x.cap
	return true
}

func (x *MyCircularDeque) DeleteLast() bool {
	if x.IsEmpty() {
		return false
	}
	x.tail++
	x.tail %= x.cap
	return true
}

func (x *MyCircularDeque) GetFront() int {
	if x.IsEmpty() {
		return -1
	}
	return x.data[(x.head-1+x.cap)%x.cap]
}

func (x *MyCircularDeque) GetRear() int {
	if x.IsEmpty() {
		return -1
	}
	return x.data[x.tail]
}

func (x *MyCircularDeque) IsEmpty() bool {
	return x.head == x.tail
}

func (x *MyCircularDeque) IsFull() bool {
	return (x.head-x.tail+x.cap)%x.cap == x.cap-1
}
