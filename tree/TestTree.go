package tree

import "fmt"

func TestBinaryTree() {
	tree := &TreeNode{&Node{Value: 1}}

	tree.Root.printValue()
	tree = &TreeNode{&Node{Value: 1, Right: &Node{Value: 3}, Left: &Node{Value: 2}}}
	fmt.Println(tree.findMaxRight())

	tree.search(4).printValue()
}
