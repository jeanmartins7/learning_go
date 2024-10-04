package linkedList

func (d *DoublyLinkedList) insertAtHead(value int) {
	temp1 := &Node{value, nil, nil}

	if d.Head == nil {
		d.Head = temp1
	} else {
		temp2 := d.Head
		d.Head = temp1
		temp1.Next = temp2
	}
	d.Length += 1
}

func (d *DoublyLinkedList) insertAtTail(value int) {
	temp1 := &Node{value, nil, nil}

	if d.Head == nil {
		d.Head = temp1
	} else {
		temp2 := d.Head
		for temp2.Next != nil {
			temp2 = temp2.Next
		}
		temp2.Next = temp1
	}
	d.Length += 1
}

func (d *DoublyLinkedList) insert(n, data int) {
	if n == 0 {
		d.insertAtHead(data)
	} else if n == d.Length-1 {
		d.insertAtTail(data)
	} else {
		temp1 := &Node{data, nil, nil}
		temp2 := d.Head

		for i := 0; i < n-1; i++ {
			temp2 = temp2.Next
		}
		temp1.Next = temp2.Next
		temp2.Next = temp1
		d.Length += 1
	}
}
