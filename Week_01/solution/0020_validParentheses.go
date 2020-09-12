package solution

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
// An input string is valid if:
//  - Open brackets must be closed by the same type of brackets.
//  - Open brackets must be closed in the correct order.

func isValid(s string) bool {
	stack := newRuneStack()
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack.push(v)
		}
		if v == ')' || v == ']' || v == '}' {
			p, ok := stack.pop()
			if !ok || !isPair(p, v) {
				return false
			}
		}
	}
	return stack.isEmpty()
}

func isPair(a, b rune) bool {
	if a == '(' {
		return b == ')'
	}
	if a == '[' {
		return b == ']'
	}
	if a == '{' {
		return b == '}'
	}
	return false
}

type runeStack struct {
	data []rune
}

func newRuneStack() *runeStack {
	return &runeStack{
		data: make([]rune, 0),
	}
}

func (x *runeStack) len() int {
	return len(x.data)
}

func (x *runeStack) isEmpty() bool {
	return x.len() == 0
}

func (x *runeStack) push(v rune) {
	x.data = append(x.data, v)
}

func (x *runeStack) pop() (rune, bool) {
	if x.isEmpty() {
		return '0', false
	}
	v := x.data[x.len()-1]
	x.data = x.data[:x.len()-1]
	return v, true
}
