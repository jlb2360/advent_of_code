package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain1(t *testing.T) {
	tests := []struct {
		pt1  []int
		pt2  []int
		want []float64
	}{
		{
			[]int{19, 13, 30, -2, 1, -2},
			[]int{18, 19, 22, -1, -1, -2},
			[]float64{14.333, 15.333},
		},
		{
			[]int{19, 13, 30, -2, 1, -2},
			[]int{20, 25, 34, -2, -2, -4},
			[]float64{11.667, 16.667},
		},
		{
			[]int{19, 13, 30, -2, 1, -2},
			[]int{12, 31, 28, -1, -2, -1},
			[]float64{6.2, 19.4},
		},
		{
			[]int{18, 19, 22, -1, -1, -2},
			[]int{20, 25, 34, -2, -2, -4},
			[]float64{math.Inf(1), math.Inf(1)},
		},
		{
			[]int{18, 19, 22, -1, -1, -2},
			[]int{20, 19, 15, 1, -5, -3},
			[]float64{math.Inf(1), math.Inf(1)},
		},
		{
			[]int{20, 25, 34, -2, -2, -4},
			[]int{20, 19, 15, 1, -5, -3},
			[]float64{math.Inf(1), math.Inf(1)},
		},
	}

	for _, test := range tests {
		ts := intercect(test.pt1, test.pt2)
		assert.Equal(t, test.want, ts)
	}

}

func TestMainPart1_2(t *testing.T) {
	f, err := os.Open("inputTest.txt")
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
				if (interc[0] >= 7) && (interc[0] <= 27) && (interc[1] >= 7) && (interc[1] <= 27) {
					counts++
				}
			}
		}
	}

	log.Println(counts)
}
