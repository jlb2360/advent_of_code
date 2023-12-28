package queue

type Node struct {
	Info []int
	Next *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (q *Queue) Push(info []int) {
	n := &Node{Info: info}
	if q.Head == nil {
		q.Head = n
		q.Tail = n
	} else {
		q.Tail.Next = n
		q.Tail = n
	}
	q.Length++
}

func (q *Queue) Pop() []int {
	if q.Head == nil {
		return nil
	}
	n := q.Head
	q.Head = q.Head.Next
	q.Length--
	return n.Info
}
