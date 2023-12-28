package queue

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (q *Queue) Enqueue(value int) {
	node := &Node{Value: value}
	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
	q.Length++
}

func (q *Queue) Dequeue() int {
	if q.Head == nil {
		return -1
	}
	value := q.Head.Value
	q.Head = q.Head.Next
	q.Length--
	return value
}
