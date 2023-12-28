package main

import (
	"bufio"
	Q "day17/queue"
	"log"
	"os"
	"strconv"
)

func main() {
	part1()
	Part2()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		ls := make([]string, 0)
		for _, c := range line {
			ls = append(ls, string(c))
		}
		grid = append(grid, ls)
	}

	q := Q.PriorityQueue{}
	q.EnQueue(Q.Tile{0, 0, 1, 0, 0}, 0)

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var minhl int
	seen := make(map[Q.Tile]bool)
	running := true
	for running {
		tile, hl, err := q.DeQueue()
		if err != nil {
			break
		}

		if tile.X == len(grid[0])-1 && tile.Y == len(grid)-1 {
			minhl = hl
			break
		}

		if _, ok := seen[tile]; ok {
			continue
		}

		seen[tile] = true

		if tile.N_steps < 3 && (tile.Dx != 0 || tile.Dy != 0) {
			nx := tile.X + tile.Dx
			ny := tile.Y + tile.Dy
			if nx >= 0 && ny >= 0 && nx < len(grid[0]) && ny < len(grid) {
				val, _ := strconv.Atoi(grid[ny][nx])
				q.EnQueue(Q.Tile{nx, ny, tile.Dx, tile.Dy, tile.N_steps + 1}, hl+val)
			}
		}

		for _, dir := range dirs {
			if dir[0] != tile.Dx && dir[1] != tile.Dy && dir[0] != -tile.Dx && dir[1] != -tile.Dy {
				nx := tile.X + dir[0]
				ny := tile.Y + dir[1]
				if nx >= 0 && ny >= 0 && nx < len(grid[0]) && ny < len(grid) {
					val, _ := strconv.Atoi(grid[ny][nx])
					q.EnQueue(Q.Tile{nx, ny, dir[0], dir[1], 1}, hl+val)
				}
			}
		}

	}

	log.Println(minhl)
}

func Part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		ls := make([]string, 0)
		for _, c := range line {
			ls = append(ls, string(c))
		}
		grid = append(grid, ls)
	}

	q := Q.PriorityQueue{}
	q.EnQueue(Q.Tile{0, 0, 1, 0, 0}, 0)

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	var minhl int
	seen := make(map[Q.Tile]bool)
	running := true
	for running {
		tile, hl, err := q.DeQueue()
		if err != nil {
			break
		}

		if tile.X == len(grid[0])-1 && tile.Y == len(grid)-1 && tile.N_steps >= 4 {
			minhl = hl
			break
		}

		if _, ok := seen[tile]; ok {
			continue
		}

		seen[tile] = true

		if tile.N_steps < 10 && (tile.Dx != 0 || tile.Dy != 0) {
			nx := tile.X + tile.Dx
			ny := tile.Y + tile.Dy
			if nx >= 0 && ny >= 0 && nx < len(grid[0]) && ny < len(grid) {
				val, _ := strconv.Atoi(grid[ny][nx])
				q.EnQueue(Q.Tile{nx, ny, tile.Dx, tile.Dy, tile.N_steps + 1}, hl+val)
			}
		}

		for _, dir := range dirs {
			if dir[0] != tile.Dx && dir[1] != tile.Dy && dir[0] != -tile.Dx && dir[1] != -tile.Dy && tile.N_steps >= 4 {
				nx := tile.X + dir[0]
				ny := tile.Y + dir[1]
				if nx >= 0 && ny >= 0 && nx < len(grid[0]) && ny < len(grid) {
					val, _ := strconv.Atoi(grid[ny][nx])
					q.EnQueue(Q.Tile{nx, ny, dir[0], dir[1], 1}, hl+val)
				}
			}
		}

	}

	log.Println(minhl)
}
