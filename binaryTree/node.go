package binaryTree

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
