package solution

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//  - push(v)  Push element v onto stack.
//  - pop()    Removes the element on top of the stack.
//  - top()    Get the top element.
//  - getMin() Retrieve the minimum element in the stack.
// Methods pop, top and getMin operations will always be called on non-empty stacks.

type MinStack struct {
	stack    *intStack
	minStack *intStack
}

func NewMinStack() MinStack {
	return MinStack{
		stack:    newIntStack(),
		minStack: newIntStack(),
	}
}

func (x *MinStack) Push(v int) {
	if x.minStack.isEmpty() || v <= x.minStack.top() {
		x.minStack.push(v)
	}
	x.stack.push(v)
}

func (x *MinStack) Pop() {
	v := x.stack.top()
	if v == x.minStack.top() {
		x.minStack.pop()
	}
	x.stack.pop()
}

func (x *MinStack) Top() int {
	return x.stack.top()
}

func (x *MinStack) GetMin() int {
	return x.minStack.top()
}
