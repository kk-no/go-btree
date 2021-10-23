package btree

import (
	"sort"
)

const (
	maxValueNum = 3
	promoteIdx = maxValueNum / 2 // Round down
)

type Node struct {
	L *Node
	R *Node

	Values []int
	Min int
	Max int
}

func NewNode() *Node {
	return &Node{Values: NewValues()}
}

func NewValues() []int {
	return make([]int, 0, maxValueNum)
}

func (n *Node) Push(value int) *Node {
	n.push(value)
	if len(n.Values) == maxValueNum {
		return n.promote()
	}
	return n
}

func (n *Node) push(value int) {
	if n.L == nil && n.R == nil {
		n.Values = append(n.Values, value)
		sort.Ints(n.Values)
		return
	}
	if value <= n.Values[0] {
		n.L.Push(value)
		return
	}
	n.R.Push(value)
}

func (n *Node) promote() *Node {
	// TODO: new promote case and append case.
	newRoot := NewNode()
	newRoot.Push(n.Values[promoteIdx])
	newRoot.L = NewNode()
	newRoot.L.Push(n.Values[promoteIdx-1])
	newRoot.R = NewNode()
	newRoot.R.Push(n.Values[promoteIdx+1])
	return newRoot
}