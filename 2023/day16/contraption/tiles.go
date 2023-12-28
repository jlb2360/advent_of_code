package contraption

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Tile struct {
	object    string
	Energized bool
}

type Contraption struct {
	tiles [][]Tile
}

func (c *Contraption) Tile(x, y int) Tile {
	return c.tiles[y][x]
}

func (c *Contraption) GrabSchematic(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, "")

		tiles := make([]Tile, len(ls))

		for j, o := range ls {
			t := Tile{object: o, Energized: false}
			tiles[j] = t
		}
		c.tiles = append(c.tiles, tiles)
		i++
	}
}

func (c *Contraption) PrintEnergized() {
	for _, row := range c.tiles {
		for _, tile := range row {
			if tile.Energized {
				fmt.Printf("%s", "#")
			} else {
				fmt.Printf("%s", ".")
			}
		}
		println()
	}
}

func (c *Contraption) PrintObjects() {
	for _, row := range c.tiles {
		for _, tile := range row {
			fmt.Printf("%s", tile.object)
		}
		println()
	}
}

func (c *Contraption) Energize(x, y int) {
	c.tiles[y][x].Energized = true
}

func (c *Contraption) AddEnergized() int {
	sum := 0
	for _, row := range c.tiles {
		for _, tile := range row {
			if tile.Energized {
				sum++
			}
		}
	}

	return sum
}

func (c *Contraption) Shape() (int, int) {
	return len(c.tiles[0]), len(c.tiles)
}
