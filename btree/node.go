package btree

import "log"

type Node struct {
	L *Node
	R *Node

	Value int
	Depth int
}

func (n *Node) Push(value int) {
	if value < n.Value {
		if n.L == nil {
			n.L = &Node{Value: value, Depth: n.Depth+1}
			log.Printf("Set %v depth %v L\n", value, n.Depth+1)
			return
		}
		n.L.Push(value)
	} else if n.Value < value {
		if n.R == nil {
			n.R = &Node{Value: value, Depth: n.Depth+1}
			log.Printf("Set %v depth %v R\n", value, n.Depth+1)
			return
		}
		n.R.Push(value)
	}
	// TODO: equal case
}