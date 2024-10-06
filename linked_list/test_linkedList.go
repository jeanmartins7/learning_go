package linkedList

func TestLinkedList() {
	list := Node{
		1, &Node{
			2, &Node{
				3, &Node{
					4, &Node{
						5, nil, nil,
					}, nil,
				}, nil,
			}, nil,
		}, nil,
	}

	nodes := &LinkedList{&list}

	nodes.insertAtEnd(6)
	nodes.insertAtEnd(7)
	nodes.removeTargetNode(7)
	nodes.addNodeInMiddle(33, 3)
	nodes.addNodeInMiddle(44, 5)
	nodes.Head.printList()
	print("Encontrado na posiçao: ", nodes.findElement(55))
}
