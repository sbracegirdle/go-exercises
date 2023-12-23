package main

import "fmt"

// Node represents a single node in the tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
	Parent *Node
}

// Tree represents the binary tree
type Tree struct {
	Root *Node
}

// Insert a new node with the given value into the tree
func (t *Tree) Insert(value int) *Tree {
	if t.Root == nil {
		t.Root = &Node{Value: value, Parent: nil}
	} else {
		t.Root.insert(value)
	}
	return t
}

// insert a new node with the given value into the tree
func (n *Node) insert(value int) {
	if n == nil {
		return
	} else if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value, Parent: n}
		} else {
			n.Left.insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value, Parent: n}
		} else {
			n.Right.insert(value)
		}
	}
}

// Print the tree in pre-order traversal
func (t *Tree) PrintTree() {
	if t.Root != nil {
		t.Root.printTree(0)
	}
	fmt.Println()
}

// printTree prints the tree in pre-order traversal
func (n *Node) printTree(level int) {
	if n == nil {
		return
	}
	format := ""
	for i := 0; i < level; i++ {
		format += "    "
	}
	format += "---[ "
	level++
	fmt.Printf(format+"%d\n", n.Value)
	n.Left.printTree(level)
	n.Right.printTree(level)
}

// Find a node with the given value
func (t *Tree) Find(value int) *Node {
	if t.Root == nil {
		return nil
	}
	return t.Root.find(value)
}

// find a node with the given value
func (n *Node) find(value int) *Node {
	if n == nil {
		return nil
	}
	switch {
	case value == n.Value:
		return n
	case value < n.Value:
		return n.Left.find(value)
	default:
		return n.Right.find(value)
	}
}

// Next finds the next higher value after the given value
func (t *Tree) Next(value int) *Node {
	node := t.Find(value)
	if node == nil {
		return nil
	}
	return node.next()
}

// next finds the next higher value after this node
func (n *Node) next() *Node {
	if n == nil {
		return nil
	}
	if n.Right != nil {
		return n.Right.minValueNode()
	}
	var succ *Node
	for p := n.Parent; p != nil; p = p.Parent {
		if p.Value > n.Value {
			succ = p
			break
		}
	}
	return succ
}

// minValueNode finds the node with the minimum value in this subtree
func (n *Node) minValueNode() *Node {
	current := n
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func main() {
	t := &Tree{}
	t.Insert(5).
		Insert(3).
		Insert(2).
		Insert(4).
		Insert(7).
		Insert(6).
		Insert(8)
	t.PrintTree() // Output: 5 3 2 4 7 6 8
}
