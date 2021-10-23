package main

import (
	"fmt"
	"github.com/kk-no/go-btree/btree"
)

// var values = []int{10, 5, 3, 2, 8, 4, 1, 6}
var values = []int{10, 5, 3, 2, 8, 4}

func main() {
	tree := btree.New()
	for _, v := range values {
		tree.Push(v)
	}
	fmt.Println(tree.Root.Values)
	fmt.Println(tree.Root.L.Values, tree.Root.R.Values)
}