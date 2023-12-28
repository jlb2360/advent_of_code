package main

import (
	"bufio"
	q "day23/queue"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func dfs(graph map[string]map[string]int, seen map[string]bool, points []int, end []int) int {
	if (points[0] == end[0]) && (points[1] == end[1]) {
		return 0
	}

	m := -9223372036854775808

	key1 := strings.Join([]string{strconv.Itoa(points[0]), strconv.Itoa(points[1])}, ",")

	seen[key1] = true
	for p, _ := range graph[key1] {
		ls := strings.Split(p, ",")
		x, _ := strconv.Atoi(ls[0])
		y, _ := strconv.Atoi(ls[1])
		point := []int{x, y}
		key2 := strings.Join([]string{strconv.Itoa(point[0]), strconv.Itoa(point[1])}, ",")
		if seen[key2] == false {
			m = max(m, dfs(graph, seen, point, end)+graph[key1][p])
		}
	}
	seen[key1] = false

	return m
}

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	points := [][]int{}
	grid := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, "")
		grid = append(grid, ls)
	}

	var start []int
	var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for x, c := range grid[0] {
		if c == "." {
			start = []int{x, 0}
			points = append(points, start)
		}
	}

	var end []int
	for x, row := range grid[len(grid)-1] {
		if row == "." {
			end = []int{x, len(grid) - 1}
			points = append(points, end)
		}
	}
	fmt.Printf("END: %v\n", end)
	fmt.Printf("START: %v\n", start)

	for y, row := range grid {
		for x, c := range row {
			if c == "#" {
				continue
			}

			neighbors := 0
			for _, dir := range directions {
				newX := x + dir[0]
				newY := y + dir[1]
				if newX < 0 || newY < 0 || newX >= len(grid) || newY >= len(grid) {
					continue
				}

				if grid[newY][newX] != "#" {
					neighbors++
				}
			}

			if neighbors >= 3 {
				points = append(points, []int{x, y})
			}

		}
	}

	fmt.Println("Points: ", len(points))

	graph := map[string]map[string]int{}

	directs := map[string][][]int{}
	directs["^"] = [][]int{{0, -1}}
	directs["v"] = [][]int{{0, 1}}
	directs[">"] = [][]int{{1, 0}}
	directs["<"] = [][]int{{-1, 0}}
	directs["."] = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, ps := range points {
		queue := q.Queue{}
		np := []int{0, ps[0], ps[1]}
		queue.Push(np)
		seen := map[string]bool{}
		key := strings.Join([]string{strconv.Itoa(np[1]), strconv.Itoa(np[2])}, ",")
		seen[key] = true

		for queue.Length > 0 {
			p := queue.Pop()

			if p[0] != 0 {
				found := false
				for _, pp := range points {
					if p[1] == pp[0] && p[2] == pp[1] {
						found = true
						break
					}
				}

				if found {
					key1 := strings.Join([]string{strconv.Itoa(p[1]), strconv.Itoa(p[2])}, ",")
					key2 := strings.Join([]string{strconv.Itoa(np[1]), strconv.Itoa(np[2])}, ",")
					if _, ok := graph[key2][key1]; !ok {
						if _, ok := graph[key2]; !ok {
							graph[key2] = map[string]int{}
						}
						graph[key2][key1] = p[0]
						continue
					}
				}
			}

			for _, dir := range directs[grid[p[2]][p[1]]] {
				nx := p[1] + dir[0]
				ny := p[2] + dir[1]
				if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid) {
					continue
				}

				if grid[ny][nx] != "#" {
					if _, ok := seen[strings.Join([]string{strconv.Itoa(nx), strconv.Itoa(ny)}, ",")]; !ok {
						queue.Push([]int{p[0] + 1, nx, ny})
						seen[strings.Join([]string{strconv.Itoa(nx), strconv.Itoa(ny)}, ",")] = true
					}
				}
			}

		}
	}

	seen := map[string]bool{}
	result := dfs(graph, seen, start, []int{len(grid[0]) - 2, len(grid) - 1})

	fmt.Printf("Result: %v\n", result)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	points := [][]int{}
	grid := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, "")
		grid = append(grid, ls)
	}

	var start []int
	var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for x, c := range grid[0] {
		if c == "." {
			start = []int{x, 0}
			points = append(points, start)
		}
	}

	var end []int
	for x, row := range grid[len(grid)-1] {
		if row == "." {
			end = []int{x, len(grid) - 1}
			points = append(points, end)
		}
	}

	for y, row := range grid {
		for x, c := range row {
			if c == "#" {
				continue
			}

			neighbors := 0
			for _, dir := range directions {
				newX := x + dir[0]
				newY := y + dir[1]
				if newX < 0 || newY < 0 || newX >= len(grid) || newY >= len(grid) {
					continue
				}

				if grid[newY][newX] != "#" {
					neighbors++
				}
			}

			if neighbors >= 3 {
				points = append(points, []int{x, y})
			}

		}
	}

	graph := map[string]map[string]int{}

	directs := map[string][][]int{}
	directs["^"] = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	directs["v"] = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	directs[">"] = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	directs["<"] = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	directs["."] = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, ps := range points {
		queue := q.Queue{}
		np := []int{0, ps[0], ps[1]}
		queue.Push(np)
		seen := map[string]bool{}
		key := strings.Join([]string{strconv.Itoa(np[1]), strconv.Itoa(np[2])}, ",")
		seen[key] = true

		for queue.Length > 0 {
			p := queue.Pop()

			if p[0] != 0 {
				found := false
				for _, pp := range points {
					if p[1] == pp[0] && p[2] == pp[1] {
						found = true
						break
					}
				}

				if found {
					key1 := strings.Join([]string{strconv.Itoa(p[1]), strconv.Itoa(p[2])}, ",")
					key2 := strings.Join([]string{strconv.Itoa(np[1]), strconv.Itoa(np[2])}, ",")
					if _, ok := graph[key2][key1]; !ok {
						if _, ok := graph[key2]; !ok {
							graph[key2] = map[string]int{}
						}
						graph[key2][key1] = p[0]
						continue
					}
				}
			}

			for _, dir := range directs[grid[p[2]][p[1]]] {
				nx := p[1] + dir[0]
				ny := p[2] + dir[1]
				if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid) {
					continue
				}

				if grid[ny][nx] != "#" {
					if _, ok := seen[strings.Join([]string{strconv.Itoa(nx), strconv.Itoa(ny)}, ",")]; !ok {
						queue.Push([]int{p[0] + 1, nx, ny})
						seen[strings.Join([]string{strconv.Itoa(nx), strconv.Itoa(ny)}, ",")] = true
					}
				}
			}

		}
	}

	seen := map[string]bool{}
	result := dfs(graph, seen, start, end)

	fmt.Printf("Result: %v\n", result)
}
