package linkedList

type DoublyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (d *DoublyLinkedList) lengthZero() {
	d.Length = 0
}
