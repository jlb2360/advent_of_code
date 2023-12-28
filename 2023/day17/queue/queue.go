package queue

import "errors"

// We are going to create a PriorityQueue, which is a queue that is sorted by priority.

type Tile struct {
	X       int //The x coordinate of the tile
	Y       int //The y coordinate of the tile
	Dx      int //The dx coordinate of the destination
	Dy      int //The dy coordinate of the destination
	N_steps int //The number of steps taken to get to this tile
}

type Node struct {
	HeatLoss int  //This will be the priority of the node
	info     Tile //This will be the information stored in the node on the for the tile and direction
	Next     *Node
}

type PriorityQueue struct {
	Head *Node
}

// Now we implement EnQueue and DeQueue functions for the PriorityQueue.

func (p *PriorityQueue) EnQueue(info Tile, heatLoss int) {
	if p.Head == nil {
		p.Head = &Node{heatLoss, info, nil}
		return
	} else if p.Head.HeatLoss > heatLoss {
		p.Head = &Node{heatLoss, info, p.Head}
		return
	} else {
		current := p.Head
		for current.Next != nil && current.Next.HeatLoss < heatLoss {
			current = current.Next
		}

		current.Next = &Node{heatLoss, info, current.Next}
	}
}

func (p *PriorityQueue) DeQueue() (Tile, int, error) {
	if p.Head == nil {
		return Tile{}, 0, errors.New("Queue is empty")
	} else {
		info := p.Head.info
		hl := p.Head.HeatLoss
		p.Head = p.Head.Next
		return info, hl, nil
	}
}
