package binaryTree

import (
	"fmt"
)

// Node is a node of the tree
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// NewNode returns a pointer to a Node instance
func newNode(val int, parent *Node) *Node {
	return &Node{
		Val:    val,
		Left:   nil,
		Right:  nil,
		Parent: parent,
	}
}

// AddNode adds a new node to the left or right of a current node.
func (n *Node) addNode(val int) {
	if n.Val < val {
		n.Right = newNode(val, n)
	} else {
		n.Left = newNode(val, n)
	}
}

func (n *Node) deleteNode() {
	if n == n.Parent.Left {
		n.Parent.Left = nil
	} else {
		n.Parent.Right = nil
	}
}

func (n *Node) branchLength() int {
	total := 1
	if n.Left != nil {
		total += n.Left.branchLength()
	}
	if n.Right != nil {
		total += n.Right.branchLength()
	}
	return total
}

func (n *Node) print() string {
	str := fmt.Sprintf("%d\n", n.Val)
	if n.Left != nil && n.Right != nil {
		str = fmt.Sprintf("%s L %s / R %s /", str, n.Left.print(), n.Right.print())
	} else if n.Left != nil {
		str = fmt.Sprintf("%s L %s", str, n.Left.print())
	} else if n.Right != nil {
		str = fmt.Sprintf("%s R %s", str, n.Right.print())
	}
	return str
}
