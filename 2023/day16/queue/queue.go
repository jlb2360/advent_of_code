package queue

import "fmt"

type Beam struct {
	X  int
	Y  int
	DX int
	DY int
}

type Node struct {
	beam Beam
	next *Node
}

type Queue struct {
	head *Node
}

// We want to insert at the end
func (l *Queue) Insert(b Beam) {
	node := &Node{beam: b, next: nil}
	if l.head == nil {
		l.head = node
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = node
	}
}

// Next we want to be able to remove from the front to perform a fifo operation
// The error tells us if we have an empty queue
func (l *Queue) Remove() (Beam, error) {
	if l.head == nil {
		return Beam{}, fmt.Errorf("Empty list")
	}
	p := l.head
	l.head = p.next
	return p.beam, nil
}
