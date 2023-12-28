package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainPart1(t *testing.T) {
	expected := 62

	f, err := os.Open("inputTest.txt")
	if err != nil {
		t.Error(err)
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

	assert.Equal(t, expected, sum)
}

func TestMainPart2(t *testing.T) {
	expected := int64(952408144115)

	f, err := os.Open("inputTest.txt")
	if err != nil {
		t.Error(err)
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

	assert.Equal(t, expected, sum)

}
