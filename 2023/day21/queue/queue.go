package queue

type Point struct {
	Object  string
	Visited bool
}

type Node struct {
	Val      int
	position []int
	Next     *Node
}

type PriorityQueue struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (q *PriorityQueue) Enqueue(val int, position []int) {
	node := &Node{Val: val, position: position}
	if q.Head == nil {
		q.Head = node
		q.Tail = node
	} else {
		if q.Head.Val > val {
			node.Next = q.Head
			q.Head = node
		} else {
			current := q.Head
			for current.Next != nil && current.Next.Val < val {
				current = current.Next
			}
			node.Next = current.Next
			current.Next = node
		}
	}
	q.Length++
}

func (q *PriorityQueue) Dequeue() ([]int, int) {
	if q.Head == nil {
		return []int{-1, -1}, -1
	}
	position := q.Head.position
	val := q.Head.Val
	q.Head = q.Head.Next
	q.Length--
	return position, val
}
