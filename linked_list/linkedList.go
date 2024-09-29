package linkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtBack(data int) {
	newNode := &Node{data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func delete() {
	fmt.Println("Hello, World!")
}

func update() {
	fmt.Println("Hello, World!")
}

func search() {
	fmt.Println("Hello, World!")
}
