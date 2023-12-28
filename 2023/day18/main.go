package main

import (
	"bufio"
	g "day18/grid"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ExecuteRight(grid [][]g.Trench, steps int, current_x *int, current_y *int, t g.Trench) [][]g.Trench {
	for i := 0; i < steps; i++ {
		*current_x++
		if *current_x >= len(grid[*current_y]) {
			grid[*current_y] = append(grid[*current_y], t)
		} else {
			grid[*current_y][*current_x] = t
		}
	}

	return grid

}

func ExecuteLeft(grid [][]g.Trench, steps int, current_x *int, current_y *int, t g.Trench) [][]g.Trench {
	for i := 0; i < steps; i++ {
		*current_x--
		if *current_x < 0 {
			temp := []g.Trench{{Dug: ".", Color: " "}}
			for i := 0; i < len(grid); i++ {
				grid[i] = append(temp, grid[i]...)
			}
			*current_x = 0
			grid[*current_y][*current_x] = t
		} else {
			grid[*current_y][*current_x] = t
		}
	}

	return grid
}

func ExecuteDown(grid [][]g.Trench, steps int, current_x *int, current_y *int, t g.Trench) [][]g.Trench {
	for i := 0; i < steps; i++ {
		*current_y++
		if *current_y >= len(grid) {
			temp := []g.Trench{}
			l := max(len(grid[0]), *current_x+1)
			for i := 0; i < l; i++ {
				temp = append(temp, g.Trench{Dug: ".", Color: " "})
			}
			grid = append(grid, temp)
			grid[*current_y][*current_x] = t
		} else {
			if *current_x >= len(grid[*current_y]) {
				for i := len(grid[*current_y]); i <= *current_x; i++ {
					grid[*current_y] = append(grid[*current_y], g.Trench{Dug: ".", Color: " "})
				}
			}
			grid[*current_y][*current_x] = t
		}
	}

	return grid
}

func ExecuteUp(grid [][]g.Trench, steps int, current_x *int, current_y *int, t g.Trench) [][]g.Trench {
	for i := 0; i < steps; i++ {
		*current_y--
		if *current_y < 0 {
			temp := [][]g.Trench{}
			temp = append(temp, []g.Trench{})
			l := max(len(grid[0]), *current_x+1)
			for i := 0; i < l; i++ {
				temp[0] = append(temp[0], g.Trench{Dug: ".", Color: " "})
			}
			grid = append(temp, grid...)
			*current_y = 0
			grid[*current_y][*current_x] = t
		} else {
			if *current_x >= len(grid[*current_y]) {
				for i := len(grid[*current_y]); i <= *current_x; i++ {
					grid[*current_y] = append(grid[*current_y], g.Trench{Dug: ".", Color: " "})
				}
			}
			grid[*current_y][*current_x] = t
		}
	}

	return grid
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

	points := [][]int{{0, 0}}

	dir := map[string][]int{"R": {0, 1}, "L": {0, -1}, "U": {-1, 0}, "D": {1, 0}}

	scanner := bufio.NewScanner(f)
	bp := 0
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, " ")
		direction := ls[0]
		steps, _ := strconv.Atoi(ls[1])

		bp += steps

		new_point := []int{0, 0}

		dirs := dir[direction]

		new_point[0] = dirs[0] * steps
		new_point[1] = dirs[1] * steps

		new_point[0] += points[len(points)-1][0]
		new_point[1] += points[len(points)-1][1]

		points = append(points, new_point)

	}
	sum := 0
	var point_y1 int
	var point_y2 int
	for y := 0; y < len(points); y++ {
		if y-1 < 0 {
			point_y1 = len(points) + y
		} else {
			point_y1 = y
		}

		if y+1 >= len(points) {
			point_y2 = (y + 1) % len(points)
		} else {
			point_y2 = y
		}

		sum += points[y][1] * (points[point_y1-1][0] - points[point_y2+1][0])
	}
	if sum < 0 {
		sum = sum * -1
	}

	sum = sum / 2
	sum = sum - bp/2 + 1

	sum = sum + bp

	fmt.Printf("Part 1: %v\n", sum)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	points := [][]int64{{0, 0}}

	dir := map[string][]int64{"R": {0, 1}, "L": {0, -1}, "U": {-1, 0}, "D": {1, 0}}

	scanner := bufio.NewScanner(f)
	var bp int64
	var direction string
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, " ")
		color := ls[2]

		steps, _ := strconv.ParseInt(color[2:len(color)-2], 16, 64)

		switch color[len(color)-2 : len(color)-1] {
		case "0":
			direction = "R"
		case "1":
			direction = "D"
		case "2":
			direction = "L"
		case "3":
			direction = "U"
		}

		bp += steps

		new_point := []int64{0, 0}

		dirs := dir[direction]

		new_point[0] = dirs[0] * steps
		new_point[1] = dirs[1] * steps

		new_point[0] += points[len(points)-1][0]
		new_point[1] += points[len(points)-1][1]

		points = append(points, new_point)

	}
	var sum int64
	var point_y1 int
	var point_y2 int
	for y := 0; y < len(points); y++ {
		if y-1 < 0 {
			point_y1 = len(points) + y
		} else {
			point_y1 = y
		}

		if y+1 >= len(points) {
			point_y2 = (y + 1) % len(points)
		} else {
			point_y2 = y
		}

		sum += points[y][1] * (points[point_y1-1][0] - points[point_y2+1][0])
	}
	if sum < 0 {
		sum = sum * -1
	}

	sum = sum / 2
	sum = sum - bp/2 + 1

	sum = sum + bp
	fmt.Printf("Part 2: %v\n", sum)
}
