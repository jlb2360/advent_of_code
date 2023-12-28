package main

import (
	"bufio"
	q "day21/queue"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainPart1(t *testing.T) {
	f, err := os.Open("inputTest.txt")
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

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := q.PriorityQueue{}
	queue.Enqueue(0, []int{currentX, currentY})

	seen := map[int][][]int{}

	iters := 0
	for queue.Length > 0 {
		position, steps := queue.Dequeue()
		iters++

		if steps > 6 {
			break
		}

		for _, direction := range directions {
			newX := position[0] + direction[0]
			newY := position[1] + direction[1]

			if newX < 0 || newY < 0 || newX >= len(garden) || newY >= len(garden) {
				continue
			}

			if garden[newY][newX].Object == "#" {
				new_s := steps + 600 // to ensure it is always behind the last step we want
				queue.Enqueue(new_s, []int{newX, newY})
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

	fmt.Printf("Iters: %d\n", iters)

	assert.Equal(t, 16, len(seen[6]))
}
