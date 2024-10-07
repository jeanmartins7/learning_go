package queue

type Queue struct {
	Front  *QueueNode
	Rear   *QueueNode
	Length int
}

type QueueNode struct {
	Data int
	Next *QueueNode
}

func (q Queue) isEmpty() bool {
	return q.Length == 0
}

func (q *Queue) enqueue(data int) {

	temp := &QueueNode{data, nil}

	if q.isEmpty() {
		q.Front = temp
	} else {
		q.Rear.Next = temp
	}

	q.Rear = temp
	q.Length++

}

func (q *Queue) dequeue() {
	if q.isEmpty() {
		println("Queue is empty")
		return
	}

	q.Length--
	q.Front = q.Front.Next

}

func (q *Queue) printQueue() {

	if q.isEmpty() {
		println("Queue is empty")
		return
	}

	current := q.Front

	for current != nil {
		println(current.Data)
		current = current.Next
	}

}
