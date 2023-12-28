package main

import (
	"bufio"
	q "day23/queue"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMainPart1(t *testing.T) {
	// ... test code ...
	f, err := os.Open("inputTest.txt")
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
	result := dfs(graph, seen, start, end)

	fmt.Printf("Result: %v\n", result)

}

func TestMainPart2(t *testing.T) {
	// ... test code ...
	f, err := os.Open("inputTest.txt")
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
