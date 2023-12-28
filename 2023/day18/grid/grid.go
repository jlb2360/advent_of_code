package grid

// list of linked lists to act as a grid
type GridNode struct {
	info LinkedList
	Next *GridNode
	Prev *GridNode
}

type Grid struct {
	Head   *GridNode
	Tail   *GridNode
	Length int
}

func (g *Grid) InsertBottom(info LinkedList) {
	if g.Head == nil {
		g.Head = &GridNode{info, nil, nil}
		g.Tail = g.Head
		g.Length++
	} else {
		g.Tail.Next = &GridNode{info, nil, g.Tail}
		g.Tail = g.Tail.Next
		g.Length++
	}
}

func (g *Grid) InsertTop(info LinkedList) {
	if g.Head == nil {
		g.Head = &GridNode{info, nil, nil}
		g.Tail = g.Head
		g.Length++
	} else {
		g.Head.Prev = &GridNode{info, g.Head, nil}
		g.Head = g.Head.Prev
		g.Length++
	}
}

func (g *Grid) Add2ListRight(info Trench, row int) {
	current := g.Head
	for i := 0; i < row; i++ {
		current = current.Next
	}
	current.info.InsertRight(info)
}

func (g *Grid) Add2ListFront(info Trench, row int) {
	current := g.Head
	for i := 0; i < row; i++ {
		current = current.Next
	}
	current.info.InsertFront(info)
}

func (g *Grid) PrintGrid() {
	current := g.Head
	for current != nil {
		current.info.PrintList()
		current = current.Next
	}
}

func (g *Grid) ExecuteRight(steps int, current_x *int, current_y *int, t Trench) {
	current := g.Head
	if current == nil {
		g.InsertBottom(LinkedList{})
		current = g.Head
	}
	for i := 0; i < *current_y; i++ {
		current = current.Next
	}
	for i := 0; i < steps; i++ {
		*current_x++
		if *current_x >= current.info.Length {
			current.info.InsertRight(t)
		} else {
			current.info.DigTrench(*current_x)
		}
	}
}

func (g *Grid) ExecuteLeft(steps int, current_x *int, current_y *int, t Trench) {
	current := g.Head
	if current == nil {
		g.InsertBottom(LinkedList{})
		current = g.Head
	}
	for i := 0; i < *current_y; i++ {
		current = current.Next
	}

	for i := 0; i < steps+1; i++ {
		*current_x--
		if *current_x < 0 {
			g.InsertColumnLeft()
			*current_x = 0
			current.info.DigTrench(*current_x)
		} else {
			current.info.DigTrench(*current_x)
		}
	}

}

func (g *Grid) ExecuteUp(steps int, current_x *int, current_y *int, t Trench) {
	current := g.Head
	if current == nil {
		g.InsertBottom(LinkedList{})
		current = g.Head
	}
	for i := 0; i < *current_y; i++ {
		current = current.Next
	}
	for i := 0; i < steps; i++ {
		*current_y--
		current = current.Prev
		if *current_y < 0 {
			g.InsertTop(LinkedList{})
			current = g.Head
			if *current_x > current.info.Length {
				temp := Trench{Dug: ".", Color: " "}
				for i := current.info.Length; i < *current_x; i++ {
					current.info.InsertFront(temp)
				}
				current.info.InsertRight(t)
			} else if *current_x < 0 {
				temp := Trench{Dug: "#", Color: " "}
				current.info.InsertFront(temp)
			} else {
				current.info.InsertAt(*current_x, t)
			}
		} else {
			if *current_x > current.info.Length {
				t := Trench{Dug: "#", Color: " "}
				for i := current.info.Length; i < *current_x; i++ {
					current.info.InsertRight(t)
				}
				current.info.InsertRight(t)
			} else if *current_x < 0 {
				temp := Trench{Dug: "#", Color: " "}
				current.info.InsertFront(temp)
			} else {
				current.info.DigTrench(*current_x)
			}
		}
	}

	if *current_y < 0 {
		*current_y = 0
	}
}

func (g *Grid) ExecuteDown(steps int, current_x *int, current_y *int, t Trench) {
	current := g.Head
	if current == nil {
		g.InsertBottom(LinkedList{})
		current = g.Head
	}
	for i := 0; i < *current_y; i++ {
		current = current.Next
	}
	for i := 0; i < steps; i++ {
		*current_y++
		if *current_y >= g.Length {
			g.InsertBottom(LinkedList{})
			current = g.Tail
			if *current_x >= current.info.Length {
				temp := Trench{Dug: ".", Color: " "}
				for i := current.info.Length; i < *current_x-1; i++ {
					current.info.InsertFront(temp)
				}
				current.info.InsertRight(t)
			} else {

				current.info.InsertAt(*current_x, t)
			}

		} else {
			if *current_x > current.info.Length {
				temp := Trench{Dug: ".", Color: " "}
				for i := current.info.Length; i < *current_x-1; i++ {
					current.info.InsertFront(temp)
				}
				current.info.InsertRight(t)
			} else {

				current.info.InsertAt(*current_x, t)
			}
			current = current.Next
		}
	}
}

func (g *Grid) InsertColumnLeft() {
	current := g.Head
	for current != nil {
		current.info.InsertFront(Trench{Dug: ".", Color: " "})
		current = current.Next
	}
}
