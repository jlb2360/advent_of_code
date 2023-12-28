package grid

import "fmt"

// Simple Linked List of Trenches
type Trench struct {
	Dug   string
	Color string
}

type Node struct {
	info Trench
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *LinkedList) InsertRight(info Trench) {
	if l.Head == nil {
		l.Head = &Node{info, nil}
		l.Tail = l.Head
		l.Length++
	} else {
		l.Tail.Next = &Node{info, nil}
		l.Tail = l.Tail.Next
		l.Length++
	}
}

// In case the trench goes negative in x
func (l *LinkedList) InsertFront(info Trench) {
	if l.Head == nil {
		l.Head = &Node{info, nil}
		l.Tail = l.Head
		l.Length++
	} else {
		l.Head = &Node{info, l.Head}
		l.Length++
	}
}

func (l *LinkedList) InsertAt(col int, info Trench) {
	current := l.Head
	for i := 0; i < col-1; i++ {
		current = current.Next
	}
	current.Next = &Node{info, current.Next}
	l.Length++
}

func (l *LinkedList) PrintList() {
	current := l.Head
	for current != nil {
		fmt.Printf("%v", current.info.Dug)
		current = current.Next
	}
	fmt.Println()
}

func (l *LinkedList) DigTrench(col int) {
	current := l.Head
	for i := 0; i < col; i++ {
		current = current.Next
	}
	current.info.Dug = "#"
}
