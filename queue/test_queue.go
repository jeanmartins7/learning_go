package queue

func TestQueue() {

	queue := Queue{}
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.dequeue()
	queue.dequeue()
	queue.printQueue()

}
