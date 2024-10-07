package stacks

import "fmt"

type StackLinkedList struct {
	Top    *StackNode
	Length int
}

type StackNode struct {
	Data int
	Next *StackNode
}

func (sl *StackLinkedList) push(data int) {

	newStackNode := &StackNode{data, sl.Top}
	sl.Top = newStackNode
	sl.Length++
}

func (sl *StackLinkedList) pop() {
	current := sl.Top

	if sl.isEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Popped element is ", sl.Length)
	sl.Length--
	sl.Top = current.Next
}

func (sl *StackLinkedList) printStack() {
	current := sl.Top

	for current != nil {
		println(current.Data)
		current = current.Next
	}
}

func (sl *StackLinkedList) isEmpty() bool {
	return sl.Length == 0
}
