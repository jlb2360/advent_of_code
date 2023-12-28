package queue

import (
	"strings"
)

type Node struct {
	Val  int
	Bid  int
	Hand string
	Next *Node
	Prev *Node
}

type PriorityQueue struct {
	Head   *Node
	Length int
}

func (q *PriorityQueue) Enqueue(hand string, bid int) {
	priority := q.getPriority(hand)
	//fmt.Printf("priority: %v\n", priority)
	newNode := &Node{Val: priority, Bid: bid, Hand: hand}
	if q.Head == nil {
		q.Head = newNode
		q.Length++
	} else {
		current := q.Head
		for current.Val < newNode.Val {
			if current.Next != nil {
				current = current.Next
			} else {
				current.Next = newNode
				newNode.Prev = current
				q.Length++
				return
			}
		}

		for current.Val >= newNode.Val {
			if current.Val > newNode.Val {
				//fmt.Printf("current.Val: %v, newNode.Val: %v\n", current.Val, newNode.Val)
				newNode.Next = current
				newNode.Prev = current.Prev
				current.Prev.Next = newNode
				current.Prev = newNode
				q.Length++
				return
			} else if current.Val == newNode.Val {
				higher := q.checkTier(current, newNode)
				if higher {
					if current.Prev != nil {
						newNode.Next = current
						newNode.Prev = current.Prev
						current.Prev.Next = newNode
						current.Prev = newNode
						q.Length++
						return
					} else {
						newNode.Next = current
						current.Prev = newNode
						q.Head = newNode
						q.Length++
						return
					}
				} else {
					if current.Next == nil {
						current.Next = newNode
						newNode.Prev = current
						q.Length++
						return
					}
					current = current.Next
				}
			}
		}
	}
}

func (q *PriorityQueue) getPriority(hand string) int {
	// finding the hand power based on the rules

	nc := make(map[string]int)
	for _, card := range hand {
		if _, ok := nc[string(card)]; !ok {
			counts := strings.Count(hand, string(card))
			nc[string(card)] = counts
		}
	}

	// 1. Royal Flush
	hp := 0
	for _, v := range nc {
		if v > hp {
			hp = v
		}
	}

	if hp == 5 {
		return 6
	} else if hp == 4 {
		return 5
	} else if hp == 3 {
		// now check for full house or three of a kind
		for _, v := range nc {
			if v == 2 {
				return 4
			}
		}
		return 3
	} else if hp == 2 {
		// now check for two pair or one pair
		np := 0
		for _, v := range nc {
			if v == 2 {
				np += 1
			}
		}
		if np == 2 {
			return 2
		} else if np == 1 {

			return 1
		}
	}
	return 0
}

func (q *PriorityQueue) checkTier(current *Node, newNode *Node) bool {
	for i := 0; i < 5; i++ {
		c1 := string(current.Hand[i])
		c2 := string(newNode.Hand[i])
		if c1 == c2 {
			continue
		} else if Compare(c1, c2) {
			return true
		} else {
			return false
		}
	}

	return false
}

func Compare(c1 string, c2 string) bool {
	if c1 == "A" {
		return true
	} else if c2 == "A" {
		return false
	} else if c1 == "K" {
		return true
	} else if c2 == "K" {
		return false
	} else if c1 == "Q" {
		return true
	} else if c2 == "Q" {
		return false
	} else if c1 == "J" {
		return true
	} else if c2 == "J" {
		return false
	} else if c1 == "T" {
		return true
	} else if c2 == "T" {
		return false
	} else {
		return c1 > c2
	}
}

func (q *PriorityQueue) Dequeue() (string, int, int) {
	if q.Head == nil {
		return "", 0, 0
	} else {
		current := q.Head
		q.Head = current.Next
		q.Length--
		return current.Hand, current.Bid, current.Val
	}
}
