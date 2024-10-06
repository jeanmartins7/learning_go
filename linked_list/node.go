package linkedList

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func (n *Node) printList() {
	current := n

	for current != nil {
		print(current.Value, "-->")
		current = current.Next
	}
}

func (n *Node) length(count int) int {
	current := n

	for current != nil {
		count++
		current = current.Next
	}

	return count
}
