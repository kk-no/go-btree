package main

import (
	"github.com/kk-no/go-btree/btree"
)

var values = []int{10, 5, 3, 2, 8, 4, 1, 6}

func main() {
	tree := btree.New()
	for _, v := range values {
		tree.Push(v)
	}
}