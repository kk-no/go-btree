package btree

type Tree struct {
	Root *Node
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) Push(value int) {
	if t.Root == nil {
		t.Root = &Node{Value: value, Depth: 1}
		return
	}
	t.Root.Push(value)
}