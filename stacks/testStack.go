package stacks

func Test_stack_push() {

	stack := StackLinkedList{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()
	stack.push(1)
	stack.printStack()

}
