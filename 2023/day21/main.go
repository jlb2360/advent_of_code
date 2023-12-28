package main

import (
	"bufio"
	q "day21/queue"
	"fmt"
	"log"
	"os"
	"strings"
)

func getCounts(garden [][]q.Point, start []int, stop int) int {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := q.PriorityQueue{}
	queue.Enqueue(0, start)

	seen := map[int][][]int{}

	iters := 0
	for queue.Length > 0 {
		position, steps := queue.Dequeue()
		iters++
		fmt.Printf("Steps: %v, Stop: %v\n", steps, stop)

		if steps > stop {
			break
		}

		for _, direction := range directions {
			newX := position[0] + direction[0]
			newY := position[1] + direction[1]

			if newX < 0 || newY < 0 || newX >= len(garden) || newY >= len(garden) {
				continue
			}

			if garden[newY][newX].Object == "#" {
				continue
			} else {
				new_s := steps + 1
				if _, ok := seen[new_s]; ok {
					// check that this is a new position
					found := false
					for _, point := range seen[new_s] {
						if point[0] == newX && point[1] == newY {
							found = true
							break
						}
					}
					if !found {
						seen[new_s] = append(seen[new_s], []int{newX, newY})
						queue.Enqueue(new_s, []int{newX, newY})
					}
				} else {
					seen[new_s] = [][]int{{newX, newY}}
					queue.Enqueue(new_s, []int{newX, newY})
				}

			}
		}
	}

	return len(seen[stop])
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

	garden := [][]q.Point{}

	var currentX int
	var currentY int
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, "")
		row := []q.Point{}

		for i, col := range ls {
			if col == "S" {
				currentX = i
				currentY = len(garden)
			}
			row = append(row, q.Point{Object: col, Visited: false})
		}

		garden = append(garden, row)
	}

	answer := getCounts(garden, []int{currentX, currentY}, 64)
	fmt.Printf("Part1: %v\n", answer)

}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	garden := [][]q.Point{}

	var currentX int
	var currentY int
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, "")
		row := []q.Point{}

		for i, col := range ls {
			if col == "S" {
				currentX = i
				currentY = len(garden)
			}
			row = append(row, q.Point{Object: col, Visited: false})
		}

		garden = append(garden, row)
	}

	size := len(garden)

	steps := 26501365
	//fmt.Printf("check_steps: %v, size/2: %v\n", check_steps, size/2)
	fmt.Printf("Start: %v, %v\n", currentX, currentY)

	grid_width := steps/size - 1

	odd := (grid_width/2*2 + 1) * (grid_width/2*2 + 1)
	even := ((grid_width + 1) / 2 * 2) * ((grid_width + 1) / 2 * 2)

	odd_points := getCounts(garden, []int{currentX, currentY}, size*2+1)
	even_points := getCounts(garden, []int{currentX, currentY}, size*2)

	top_corner := getCounts(garden, []int{currentX, size - 1}, size-1)
	right_corner := getCounts(garden, []int{0, currentY}, size-1)
	bottom_corner := getCounts(garden, []int{currentX, 0}, size-1)
	left_corner := getCounts(garden, []int{size - 1, currentY}, size-1)

	topright_small := getCounts(garden, []int{0, size - 1}, size/2-1)
	topleft_small := getCounts(garden, []int{size - 1, size - 1}, size/2-1)
	bottomright_small := getCounts(garden, []int{0, 0}, size/2-1)
	bottomleft_small := getCounts(garden, []int{size - 1, 0}, size/2-1)

	topright_large := getCounts(garden, []int{0, size - 1}, (size*3)/2-1)
	topleft_large := getCounts(garden, []int{size - 1, size - 1}, (size*3)/2-1)
	bottomright_large := getCounts(garden, []int{0, 0}, (size*3)/2-1)
	bottomleft_large := getCounts(garden, []int{size - 1, 0}, (size*3)/2-1)

	sum := odd*odd_points + even*even_points
	sum += top_corner + right_corner + bottom_corner + left_corner
	sum += (topright_small + topleft_small + bottomright_small + bottomleft_small) * (grid_width + 1)
	sum += (topright_large + topleft_large + bottomright_large + bottomleft_large) * (grid_width)
	fmt.Printf("Part2: %v\n", sum)
}
