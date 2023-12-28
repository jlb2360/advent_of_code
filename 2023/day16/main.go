package main

import (
	"bufio"
	Q "day16/queue"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	i := 0

	matrix := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, "")

		tiles := make([]string, len(ls))

		for j, o := range ls {
			tiles[j] = o
		}
		matrix = append(matrix, tiles)
		i++
	}

	fmt.Printf("Length of matrix: %v\n", len(matrix))
	q := Q.Queue{}

	//starting position
	q.Insert(Q.Beam{X: -1, Y: 0, DX: 1, DY: 0})
	running := true

	seen := make(map[string]bool)
	for running {
		beam, err := q.Remove()
		if err != nil {
			running = false
			break
		}

		fmt.Printf("x: %v, y: %v, dx: %v, dy: %v\n", beam.X, beam.Y, beam.DX, beam.DY)

		beam.X += beam.DX
		beam.Y += beam.DY

		//Check if beam is out of bounds
		if beam.X < 0 || beam.Y < 0 || beam.X >= len(matrix[0]) || beam.Y >= len(matrix) {
			continue
		}

		if (matrix[beam.Y][beam.X] == ".") || (matrix[beam.Y][beam.X] == "-" && beam.DX != 0) || (matrix[beam.Y][beam.X] == "|" && beam.DY != 0) {
			label := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(beam.DX) + "," + strconv.Itoa(beam.DY)
			if _, ok := seen[label]; !ok {
				seen[label] = true
				q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: beam.DX, DY: beam.DY})
			}
		} else if matrix[beam.Y][beam.X] == "/" {
			// we want to flip the positions of dx and dy and negate them
			x, y, dx, dy := beam.X, beam.Y, -beam.DY, -beam.DX
			label := strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(dx) + "," + strconv.Itoa(dy)
			if _, ok := seen[label]; !ok {
				seen[label] = true
				q.Insert(Q.Beam{X: x, Y: y, DX: dx, DY: dy})
			}
		} else if matrix[beam.Y][beam.X] == "\\" {
			// we want to flip the positions of dx and dy
			x, y, dx, dy := beam.X, beam.Y, beam.DY, beam.DX
			label := strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(dx) + "," + strconv.Itoa(dy)
			if _, ok := seen[label]; !ok {
				seen[label] = true
				q.Insert(Q.Beam{X: x, Y: y, DX: dx, DY: dy})
			}

		} else if matrix[beam.Y][beam.X] == "-" {
			dx1, dy1 := 1, 0
			dx2, dy2 := -1, 0

			label1 := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(dx1) + "," + strconv.Itoa(dy1)
			label2 := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(dx2) + "," + strconv.Itoa(dy2)
			if _, ok := seen[label1]; !ok {
				seen[label1] = true
				q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: dx1, DY: dy1})
			}
			if _, ok := seen[label2]; !ok {
				seen[label2] = true
				q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: dx2, DY: dy2})
			}
		} else if matrix[beam.Y][beam.X] == "|" {
			dx1, dy1 := 0, 1
			dx2, dy2 := 0, -1

			label1 := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(dx1) + "," + strconv.Itoa(dy1)
			label2 := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(dx2) + "," + strconv.Itoa(dy2)

			if _, ok := seen[label1]; !ok {
				seen[label1] = true
				q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: dx1, DY: dy1})
			}

			if _, ok := seen[label2]; !ok {
				seen[label2] = true
				q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: dx2, DY: dy2})
			}
		}
	}
	s := 0

	sum_map := map[string]bool{}

	for k, v := range seen {
		print(k, v, "\n")
		ls := strings.Split(k, ",")
		label := ls[0] + "," + ls[1]
		if v {
			sum_map[label] = true
		}
	}

	s = len(sum_map)

	fmt.Printf("Part 1: %d\n", s)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0

	matrix := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, "")

		tiles := make([]string, len(ls))

		for j, o := range ls {
			tiles[j] = o
		}
		matrix = append(matrix, tiles)
		i++
	}

	m := 0

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if y == 0 || y == len(matrix)-1 || x == 0 || x == len(matrix[0])-1 {
				q := Q.Queue{}
				seen := make(map[string]bool)
				if (y == 0) && (x != 0 || x != len(matrix[0])-1) {
					q.Insert(Q.Beam{X: x, Y: y - 1, DX: 0, DY: 1})
				} else if (y == len(matrix)-1) && (x != 0 || x != len(matrix[0])-1) {
					q.Insert(Q.Beam{X: x, Y: y + 1, DX: 0, DY: -1})
				} else if (x == 0) && (y != 0 || y != len(matrix)-1) {
					q.Insert(Q.Beam{X: x - 1, Y: y, DX: 1, DY: 0})
				} else if (x == len(matrix[0])-1) && (y != 0 || y != len(matrix)-1) {
					q.Insert(Q.Beam{X: x + 1, Y: y, DX: -1, DY: 0})
				}

				//starting position
				running := true

				for running {
					beam, err := q.Remove()
					if err != nil {
						running = false
						break
					}

					beam.X += beam.DX
					beam.Y += beam.DY

					//Check if beam is out of bounds
					if beam.X < 0 || beam.Y < 0 || beam.X >= len(matrix[0]) || beam.Y >= len(matrix) {
						continue
					}

					if (matrix[beam.Y][beam.X] == ".") || (matrix[beam.Y][beam.X] == "-" && beam.DX != 0) || (matrix[beam.Y][beam.X] == "|" && beam.DY != 0) {
						label := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(beam.DX) + "," + strconv.Itoa(beam.DY)
						if _, ok := seen[label]; !ok {
							seen[label] = true
							q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: beam.DX, DY: beam.DY})
						}
					} else if matrix[beam.Y][beam.X] == "/" {
						// we want to flip the positions of dx and dy and negate them
						x, y, dx, dy := beam.X, beam.Y, -beam.DY, -beam.DX
						label := strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(dx) + "," + strconv.Itoa(dy)
						if _, ok := seen[label]; !ok {
							seen[label] = true
							q.Insert(Q.Beam{X: x, Y: y, DX: dx, DY: dy})
						}
					} else if matrix[beam.Y][beam.X] == "\\" {
						// we want to flip the positions of dx and dy
						x, y, dx, dy := beam.X, beam.Y, beam.DY, beam.DX
						label := strconv.Itoa(x) + "," + strconv.Itoa(y) + "," + strconv.Itoa(dx) + "," + strconv.Itoa(dy)
						if _, ok := seen[label]; !ok {
							seen[label] = true
							q.Insert(Q.Beam{X: x, Y: y, DX: dx, DY: dy})
						}

					} else if matrix[beam.Y][beam.X] == "-" {
						dx1, dy1 := 1, 0
						dx2, dy2 := -1, 0

						label1 := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(dx1) + "," + strconv.Itoa(dy1)
						label2 := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(dx2) + "," + strconv.Itoa(dy2)
						if _, ok := seen[label1]; !ok {
							seen[label1] = true
							q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: dx1, DY: dy1})
						}
						if _, ok := seen[label2]; !ok {
							seen[label2] = true
							q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: dx2, DY: dy2})
						}
					} else if matrix[beam.Y][beam.X] == "|" {
						dx1, dy1 := 0, 1
						dx2, dy2 := 0, -1

						label1 := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(dx1) + "," + strconv.Itoa(dy1)
						label2 := strconv.Itoa(beam.X) + "," + strconv.Itoa(beam.Y) + "," + strconv.Itoa(dx2) + "," + strconv.Itoa(dy2)

						if _, ok := seen[label1]; !ok {
							seen[label1] = true
							q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: dx1, DY: dy1})
						}

						if _, ok := seen[label2]; !ok {
							seen[label2] = true
							q.Insert(Q.Beam{X: beam.X, Y: beam.Y, DX: dx2, DY: dy2})
						}
					}
				}

				s := 0

				sum_map := map[string]bool{}

				for k, v := range seen {
					ls := strings.Split(k, ",")
					label := ls[0] + "," + ls[1]
					if v {
						sum_map[label] = true
					}
				}

				s = len(sum_map)

				if m < s {
					m = s
				}

			}
		}
	}

	fmt.Printf("Part 2: %d\n", m)

}
