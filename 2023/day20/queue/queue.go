package queue

type Node struct {
	Destination string
	Signal      string
	Sender      string
	Next        *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (q *Queue) Enqueue(destination string, signal string, sender string) {
	node := &Node{Destination: destination, Signal: signal, Sender: sender}

	if q.Head == nil {
		q.Head = node
		q.Tail = node
		q.Length = 1
	} else {
		q.Tail.Next = node
		q.Tail = node
		q.Length++
	}
}

func (q *Queue) Dequeue() (string, string, string) {
	if q.Head == nil {
		return "", "", ""
	}

	destination := q.Head.Destination
	signal := q.Head.Signal
	sender := q.Head.Sender

	q.Head = q.Head.Next
	q.Length--

	return destination, signal, sender
}
