package binaryTree

import (
	"errors"
	"log"
)

// Tree is a struct for a binary tree
type Tree struct {
	Root     *Node
	current  *Node
	sCurrent *Node
}

// New returns a pointer to a tree instance
func New(vals []int) *Tree {
	newTree := &Tree{
		Root:     newNode(vals[0], nil),
		current:  nil,
		sCurrent: nil,
	}
	newTree.current = newTree.Root
	newTree.sCurrent = newTree.Root
	for i := 1; i < len(vals); i++ {
		newTree.Add(vals[i])
	}
	return newTree
}

// Add adds a Node with a value
func (t *Tree) Add(val int) {
	if val < t.current.Val {
		t.current.Left = newNode(val, t.current)
	} else {
		t.current.Right = newNode(val, t.current)
	}
}

// Search finds a value in the tree
func (t *Tree) Search(val int) (*Node, error) {
	if t.sCurrent != nil {
		if t.sCurrent.Val != val {
			if t.sCurrent.Val > val {
				t.sCurrent = t.sCurrent.Left
				return t.Search(val)
			}
			t.sCurrent = t.sCurrent.Right
			return t.Search(val)
		}
		res := t.sCurrent
		t.sCurrent = t.Root
		return res, nil
	}
	t.sCurrent = t.Root
	return nil, errors.New("Could not find the requested value")
}

// Print prints the tree
func (t *Tree) Print() {
	log.Println(t.Root.print())
}
