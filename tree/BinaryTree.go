package tree

import "fmt"

type Node struct {
	Value int
	Right *Node
	Left  *Node
}

type TreeNode struct {
	Root *Node
}

func (n *Node) printValue() {
	if n == nil {
		fmt.Println("nó vazio")
	} else {
		fmt.Println("valor do nó: ", n.Value)
	}
}

func (bt *TreeNode) findMaxRight() *Node {
	current := bt.Root

	for current.Right != nil {
		current = current.Right
	}

	return current
}

func (bt *TreeNode) findMaxLeft() *Node {
	current := bt.Root

	for current.Left != nil {
		current = current.Left
	}

	return current
}

func (bt *TreeNode) search(data int) *Node {
	if bt.Root == nil {
		return nil
	}
	current := bt.Root

	for current != nil {
		if data == current.Value {
			return current
		}

		if data < current.Value {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}
