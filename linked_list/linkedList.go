package linkedList

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) insertAtBeginning(data int) {
	newNode := &Node{Value: data, Next: ll.Head}
	ll.Head = newNode
}

func (ll *LinkedList) insertAtEnd(data int) {
	newNode := &Node{Value: data}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

func (ll *LinkedList) printListLinked() {
	current := ll.Head

	for current != nil {
		println(current.Value)
		current = current.Next
	}
}

func (ll *LinkedList) removeFirstNode() {
	if ll.Head == nil {
		return
	}

	ll.Head = ll.Head.Next
	return
}

func (ll *LinkedList) removeLastNode() {

	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	previous := ll.Head

	for previous.Next.Next != nil {
		previous = previous.Next
	}

	previous.Next = nil
	return
}

func (ll *LinkedList) removeTargetNode(target int) {

	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	previous := ll.Head

	if previous.Value == target {
		ll.removeFirstNode()
		return
	}

	for previous.Next.Next != nil {

		if previous.Next.Value == target {
			if previous.Next.Next == nil {
				previous.Next = nil
				return
			}
			previous.Next = previous.Next.Next
			return
		}

		previous = previous.Next
		if previous.Next.Next == nil && previous.Next.Value == target {
			ll.removeLastNode()
			return
		}
	}

	return

}

func (ll *LinkedList) addNodeInMiddle(data int, position int) {

	newNode := &Node{Value: data}

	if position == 1 {
		newNode.Next = ll.Head
		ll.Head = newNode
		return
	}

	previous := ll.Head
	count := 1

	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	newNode.Next = current
	previous.Next = newNode

}

func (ll *LinkedList) removeNodeGivenPosition(position int) {
	if position <= 0 {
		return
	}

	if position == 1 {
		ll.Head = ll.Head.Next
		return
	}

	count := 1
	previous := ll.Head

	for count < position-1 {
		previous = previous.Next
		count++
	}

	current := previous.Next
	previous.Next = current.Next
	current = nil

}

func (ll *LinkedList) findElement(target int) int {

	current := ll.Head
	count := 1

	for current != nil {
		if target == current.Value {
			return count
		}

		count++
		current = current.Next

	}

	return -1
}
