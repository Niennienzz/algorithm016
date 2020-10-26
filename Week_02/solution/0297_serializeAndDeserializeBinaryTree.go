package solution

import (
	"strconv"
	"strings"
)

type Codec struct {
	sep   byte
	empty byte
}

func NewCodec() Codec {
	return Codec{
		sep:   ',',
		empty: '#',
	}
}

func (x *Codec) Serialize(root *TreeNode) string {
	sb := new(strings.Builder)
	x.serializeHelper(root, sb)
	return sb.String()
}

func (x *Codec) serializeHelper(node *TreeNode, sb *strings.Builder) {
	if node == nil {
		sb.WriteByte(x.empty)
		sb.WriteByte(x.sep)
	} else {
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(x.sep)
		x.serializeHelper(node.Left, sb)
		x.serializeHelper(node.Right, sb)
	}
}

func (x *Codec) Deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	splits := strings.Split(data, string(x.sep))
	return x.deserializeHelper(newStringQueue(splits))
}

func (x *Codec) deserializeHelper(q *stringQueue) *TreeNode {
	if q.len() == 0 {
		return nil
	}

	s := q.pop()
	if s == string(x.empty) {
		return nil
	}

	val, _ := strconv.Atoi(s)
	node := &TreeNode{Val: val}

	node.Left = x.deserializeHelper(q)
	node.Right = x.deserializeHelper(q)

	return node
}
