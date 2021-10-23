package btree

type Tree struct {
	Root *Node
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) Push(value int) {
	if t.Root == nil {
		t.Root = NewNode()
	}
	t.Root = t.Root.Push(value)
}