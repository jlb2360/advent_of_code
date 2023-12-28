package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// calculate intercept
func intercect(pt1 []int, pt2 []int) []float64 {
	// pt = [x1, y1, z1, dx, dy, dz]
	m1 := float64(pt1[4]) / float64(pt1[3])
	m2 := float64(pt2[4]) / float64(pt2[3])

	if m1 == m2 {
		// parallel lines
		return []float64{math.Inf(1), math.Inf(1)}
	}

	b1 := float64(pt1[1]) - m1*float64(pt1[0])
	b2 := float64(pt2[1]) - m2*float64(pt2[0])

	i_x := (b1 - b2) / (m2 - m1)
	i_y := m1*i_x + b1

	// round i_x and i_y to 3 decimal places
	i_x = roundFloat(i_x, 3)
	i_y = roundFloat(i_y, 3)

	if (pt1[4] > 0) && (i_y < float64(pt1[1])) {
		// happened in the past
		return []float64{math.Inf(1), math.Inf(1)}
	} else if (pt1[4] < 0) && (i_y > float64(pt1[1])) {
		// happened in the past
		return []float64{math.Inf(1), math.Inf(1)}
	}

	if (pt2[4] > 0) && (i_y < float64(pt2[1])) {
		// happened in the past
		return []float64{math.Inf(1), math.Inf(1)}
	} else if (pt2[4] < 0) && (i_y > float64(pt2[1])) {
		// happened in the past
		return []float64{math.Inf(1), math.Inf(1)}
	}

	return []float64{i_x, i_y}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	pts := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "@", ",", -1)
		ls := strings.Split(line, ",")
		pt := []int{}
		for _, c := range ls {
			c = strings.TrimSpace(c)
			c, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			pt = append(pt, c)
		}
		pts = append(pts, pt)
	}

	counts := 0
	for i := 0; i < len(pts); i++ {
		for j := i + 1; j < len(pts); j++ {
			interc := intercect(pts[i], pts[j])
			if interc[0] != math.Inf(1) {
				if (interc[0] >= 200000000000000) && (interc[0] <= 400000000000000) && (interc[1] >= 200000000000000) && (interc[1] <= 400000000000000) {
					counts++
				}
			}
		}
	}

	log.Println(counts)
}
