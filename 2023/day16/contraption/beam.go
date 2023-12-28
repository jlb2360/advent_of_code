package contraption

import "fmt"

type Beam struct {
	posx      int
	posy      int
	direction string
	path      map[int]map[int][]string
}

type Node struct {
	info Beam
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	Length int
}

func (l *LinkedList) Insert(b Beam) {
	node := &Node{info: b, next: nil}
	if l.head == nil {
		l.head = node
		l.tail = node
		l.Length++
	} else {
		p := l.tail
		p.next = node
		l.tail = node
		l.Length++
	}
}

func (b *Beam) MoveDot() {
	switch b.direction {
	case "right":
		b.posx++
	case "left":
		b.posx--
	case "up":
		b.posy--
	case "down":
		b.posy++
	}

}

func (b *Beam) MoveNorthMirror() {
	switch b.direction {
	case "right":
		{
			b.direction = "up"
			b.posy--
		}
	case "left":
		{
			b.direction = "down"
			b.posy++
		}
	case "up":
		{
			b.direction = "right"
			b.posx++
		}
	case "down":
		{
			b.direction = "left"
			b.posx--
		}
	}
}

func (b *Beam) MoveSouthMirror() {
	switch b.direction {
	case "right":
		{
			b.direction = "down"
			b.posy++
		}
	case "left":
		{
			b.direction = "up"
			b.posy--
		}
	case "up":
		{
			b.direction = "left"
			b.posx--
		}
	case "down":
		{
			b.direction = "right"
			b.posx++
		}
	}
}

func (b *Beam) MoveDash(l *LinkedList) {
	switch b.direction {
	case "right":
		b.posx += 1
	case "left":
		b.posx -= 1
	case "up":
		{
			b.direction = "right"
			if b.posx-1 >= 0 {
				begining_path := make(map[int]map[int][]string)
				for k, v := range b.path {
					begining_path[k] = make(map[int][]string)
					for k2, v2 := range v {
						begining_path[k][k2] = v2
					}
				}
				if _, ok := begining_path[b.posx-1]; !ok {
					begining_path[b.posx-1] = make(map[int][]string)
				}
				begining_path[b.posx-1][b.posy] = append(begining_path[b.posx-1][b.posy], "left")
				l.Insert(Beam{posx: b.posx - 1, posy: b.posy, direction: "left", path: begining_path})
			}
			b.posx += 1
		}
	case "down":
		{
			b.direction = "right"
			if b.posx-1 >= 0 {
				begining_path := make(map[int]map[int][]string)

				for k, v := range b.path {
					begining_path[k] = make(map[int][]string)
					for k2, v2 := range v {
						begining_path[k][k2] = v2
					}
				}
				if _, ok := begining_path[b.posx-1]; !ok {
					begining_path[b.posx-1] = make(map[int][]string)
				}
				begining_path[b.posx-1][b.posy] = append(begining_path[b.posx-1][b.posy], "left")
				l.Insert(Beam{posx: b.posx - 1, posy: b.posy, direction: "left", path: begining_path})
			}
			b.posx += 1
		}
	}
}

func (b *Beam) MoveOr(l *LinkedList) {
	switch b.direction {
	case "right":
		{
			b.direction = "down"
			if b.posy-1 >= 0 {
				begining_path := make(map[int]map[int][]string)
				for k, v := range b.path {
					begining_path[k] = make(map[int][]string)
					for k2, v2 := range v {
						begining_path[k][k2] = v2
					}
				}

				if _, ok := begining_path[b.posx]; !ok {
					begining_path[b.posx] = make(map[int][]string)
				}

				begining_path[b.posx][b.posy-1] = append(begining_path[b.posx][b.posy-1], "up")
				l.Insert(Beam{posx: b.posx, posy: b.posy - 1, direction: "up", path: begining_path})
			}
			b.posy += 1
		}
	case "left":
		{
			b.direction = "down"
			if b.posy-1 >= 0 {
				begining_path := make(map[int]map[int][]string)
				for k, v := range b.path {
					begining_path[k] = make(map[int][]string)
					for k2, v2 := range v {
						begining_path[k][k2] = v2
					}
				}
				if _, ok := begining_path[b.posx]; !ok {
					begining_path[b.posx] = make(map[int][]string)
				}
				begining_path[b.posx][b.posy-1] = append(begining_path[b.posx][b.posy-1], "up")
				l.Insert(Beam{posx: b.posx, posy: b.posy - 1, direction: "up", path: begining_path})
			}
			b.posy += 1
		}
	case "up":
		b.posy -= 1
	case "down":
		b.posy += 1
	}
}

func (l *LinkedList) delete(i int) {

	if i == 0 {
		if l.head.next != nil {
			l.head = l.head.next
			l.Length--
			return
		} else {
			l.head = nil
			l.tail = nil
			l.Length--
			return
		}
	} else {
		p := l.head
		for j := 0; j < i-1; j++ {
			p = p.next
		}
		if p.next.next != nil {
			p.next = p.next.next
			l.Length--
			return
		} else {
			p.next = nil
			l.tail = p
			l.Length--
			return
		}
	}
}

func (l *LinkedList) Cycle(c Contraption) {
	beamNode := l.head
	beam_num := 0
	for i := 0; i < l.Length; i++ {
		deleted := false
		beam_pos := &beamNode.info
		tile := c.Tile(beam_pos.posx, beam_pos.posy).object
		c.Energize(beam_pos.posx, beam_pos.posy)
		if tile == "." {
			beam_pos.MoveDot()
			if (beam_pos.posx < 0 || beam_pos.posy < 0) || (beam_pos.posx >= len(c.tiles[0]) || beam_pos.posy >= len(c.tiles)) {
				l.delete(beam_num)
				deleted = true

			} else {
				c.Energize(beam_pos.posx, beam_pos.posy)
			}
		} else if tile == "/" {
			beam_pos.MoveNorthMirror()
			if (beam_pos.posx < 0 || beam_pos.posy < 0) || (beam_pos.posx >= len(c.tiles[0]) || beam_pos.posy >= len(c.tiles)) {
				l.delete(beam_num)
				deleted = true
			} else {
				c.Energize(beam_pos.posx, beam_pos.posy)
			}
		} else if tile == "\\" {
			beam_pos.MoveSouthMirror()
			if (beam_pos.posx < 0 || beam_pos.posy < 0) || (beam_pos.posx >= len(c.tiles[0]) || beam_pos.posy >= len(c.tiles)) {
				l.delete(beam_num)
				deleted = true
			} else {
				c.Energize(beam_pos.posx, beam_pos.posy)
			}
		} else if tile == "-" {
			beam_pos.MoveDash(l)
			if (beam_pos.posx < 0 || beam_pos.posy < 0) || (beam_pos.posx >= len(c.tiles[0]) || beam_pos.posy >= len(c.tiles)) {
				l.delete(beam_num)
				deleted = true
			} else {
				c.Energize(beam_pos.posx, beam_pos.posy)
			}
		} else if tile == "|" {
			beam_pos.MoveOr(l)
			if (beam_pos.posx < 0 || beam_pos.posy < 0) || (beam_pos.posx >= len(c.tiles[0]) || beam_pos.posy >= len(c.tiles)) {
				l.delete(beam_num)
				deleted = true
			} else {
				c.Energize(beam_pos.posx, beam_pos.posy)
			}
		}

		if deleted == false {
			if _, ok := beam_pos.path[beam_pos.posx]; ok {
				if _, ok := beam_pos.path[beam_pos.posx][beam_pos.posy]; ok {
					found := false
					for _, dir := range beam_pos.path[beam_pos.posx][beam_pos.posy] {
						if (dir == beam_pos.direction) && !found {

							l.delete(beam_num)
							found = true
						}
					}
					if !found {
						beam_pos.path[beam_pos.posx][beam_pos.posy] = append(beam_pos.path[beam_pos.posx][beam_pos.posy], beam_pos.direction)
					}
				} else {
					beam_pos.path[beam_pos.posx][beam_pos.posy] = append(beam_pos.path[beam_pos.posx][beam_pos.posy], beam_pos.direction)
				}
			} else {
				beam_pos.path[beam_pos.posx] = make(map[int][]string)
				beam_pos.path[beam_pos.posx][beam_pos.posy] = append(beam_pos.path[beam_pos.posx][beam_pos.posy], beam_pos.direction)
			}
		}

		l.CheckBeams(c)

		beam_num++
		beamNode = beamNode.next

	}
}

func (l *LinkedList) Run(c Contraption, start_x int, start_y int, start_direction string) {
	begining_path := make(map[int]map[int][]string)
	begining_path[start_x] = make(map[int][]string)
	begining_path[start_x][start_y] = append(begining_path[start_x][start_y], start_direction)

	l.Insert(Beam{posx: start_x, posy: start_y, direction: start_direction, path: begining_path})
	for l.Length > 0 {
		l.Cycle(c)
	}
}

func (l *LinkedList) Print() {
	p := l.head
	i := 0
	for p != nil {

		fmt.Printf("%v: (%v, %v) %v -> ", i, p.info.posx, p.info.posy, p.info.direction)
		i++
		p = p.next
	}

	fmt.Println()
}

func (l *LinkedList) CheckBeams(c Contraption) {
	p := l.head
	beam_num := 0
	for p != nil {
		if p.info.posx < 0 || p.info.posy < 0 || p.info.posx >= len(c.tiles[0]) || p.info.posy >= len(c.tiles) {
			l.delete(beam_num)
		}
		beam_num++
		p = p.next
	}
}
