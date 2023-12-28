package main

import (
	"fmt"
	"reflect"
)

type Platform struct {
	Platform [][]string
}

func (b *Platform) GetPlatform() [][]string {
	return b.Platform
}

func (b *Platform) SetPlatform(p [][]string) {
	plat := Platform{Platform: make([][]string, len(p))}
	for i, row := range p {
		plat.Platform[i] = make([]string, len(row))
		for j, col := range row {
			plat.Platform[i][j] = col
		}
	}
}

func (b *Platform) Copy() Platform {
	plat := Platform{Platform: make([][]string, len(b.Platform))}

	for i, row := range b.Platform {
		plat.Platform[i] = make([]string, len(row))
		for j, col := range row {
			plat.Platform[i][j] = col
		}
	}

	return plat
}

func (b *Platform) PrintPlat() {
	for _, row := range b.Platform {
		for _, col := range row {
			fmt.Printf("%v", col)
		}
		fmt.Printf("\n")
	}
}

func (b *Platform) CalcLoad() int {
	load := 0
	for i, row := range b.Platform {
		for _, col := range row {
			if col == "O" {
				load += len(b.Platform) - i
			}

		}
	}

	return load
}

func (b *Platform) MoveUp(row int, col int) {
	if (row-1 >= 0) && (b.Platform[row-1][col] == ".") {
		b.MoveUp(row-1, col)
	} else {
		b.Platform[row][col] = "O"
	}
}

func (b *Platform) MoveRight(row int, col int) {
	if (col+1 < len(b.Platform[row])) && (b.Platform[row][col+1] == ".") {
		b.MoveRight(row, col+1)
	} else {
		b.Platform[row][col] = "0"
	}
}

func (b *Platform) MoveLeft(row int, col int) {
	if (col-1 >= 0) && (b.Platform[row][col-1] == ".") {
		b.MoveLeft(row, col-1)
	} else {
		b.Platform[row][col] = "0"
	}
}

func (b *Platform) MoveDown(row int, col int) {
	if (row+1 < len(b.Platform)) && (b.Platform[row+1][col] == ".") {
		b.MoveDown(row+1, col)
	} else {
		b.Platform[row][col] = "0"
	}
}

func (b *Platform) rotate90() {
	new := make([][]string, len(b.Platform[0]))

	for i := 0; i < len(b.Platform[0]); i++ {
		new[i] = make([]string, len(b.Platform))
		for j := 0; j < len(b.Platform); j++ {
			new[i][j] = b.Platform[len(b.Platform)-j-1][i]
		}
	}

	b.Platform = new
}

func (b *Platform) cycleup() {
	for i, row := range b.Platform {
		for j, col := range row {
			if (col == "O") || (col == "0") {
				b.Platform[i][j] = "."
				b.MoveUp(i, j)
			}
		}
	}
}

func (b *Platform) Cycle() {
	for i := 0; i < 4; i++ {
		b.cycleup()
		b.rotate90()
	}

}

func (b *Platform) SumRC() int {
	sum := 0
	for i, row := range b.Platform {
		for j, col := range row {
			if (col == "O") || (col == "0") {
				sum += j * (i + 1)
			}
		}
	}
	return sum
}

func sumrc(p [][]string) int {
	sum := 0
	for i, row := range p {
		for j, col := range row {
			if (col == "O") || (col == "0") {
				sum += j * (i + 1)
			}
		}
	}
	return sum
}

func (b *Platform) CycleN(n int) int {

	mapPlats := map[int][][]string{}
	keys := []int{}
	loads := []int{}

	mapPlats[b.SumRC()] = b.Platform
	keys = append(keys, b.SumRC())
	loads = append(loads, b.CalcLoad())

	for i := 0; i < n; i++ {
		b.Cycle()
		if _, ok := mapPlats[b.SumRC()]; ok {
			iter := i + 1
			var offset int
			for j := 0; j < len(keys); j++ {
				if keys[j] == b.SumRC() {
					offset = j
				}
			}

			state, ok := mapPlats[keys[((n-iter)%(iter-offset))+offset]]
			load := loads[((n-iter)%(iter-offset))+offset]
			if ok {

				b.Platform = state
			}
			return load
		} else {
			mapPlats[b.SumRC()] = b.GetPlatform()
			keys = append(keys, b.SumRC())
			loads = append(loads, b.CalcLoad())
		}
	}

	return b.CalcLoad()
}

func (b *Platform) MoveAll() {
	for i, row := range b.Platform {
		for j, col := range row {
			if (col == "O") || (col == "0") {
				b.Platform[i][j] = "."
				b.MoveUp(i, j)
			}
		}
	}
}

func ComparePlat(p1 Platform, p2 Platform) bool {
	if reflect.DeepEqual(p1.Platform, p2.Platform) {
		return true
	}

	return false
}

func CheckPlats(ps []Platform, p Platform) bool {
	for _, plat := range ps {
		if ComparePlat(plat, p) {
			return true
		}
	}

	return false
}
