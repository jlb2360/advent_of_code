package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainPart1(t *testing.T) {
	// your test code here

	expected := 19114

	f, err := os.Open("inputTest.txt")

	if err != nil {
		t.Errorf("Error reading input file: %s", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	worflow := map[string][]string{}
	parts := []Part{}

	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, "{")

		if len(ls[0]) > 1 {
			nls := strings.Split(ls[1][:len(ls[1])-1], ",")
			worflow[ls[0]] = nls
		} else {
			if line != "" {
				nls := strings.Split(ls[1][:len(ls[1])-1], ",")
				x, _ := strconv.Atoi(strings.Split(nls[0], "=")[1])
				m, _ := strconv.Atoi(strings.Split(nls[1], "=")[1])
				a, _ := strconv.Atoi(strings.Split(nls[2], "=")[1])
				s, _ := strconv.Atoi(strings.Split(nls[3], "=")[1])
				parts = append(parts, Part{x, m, a, s})
			}
		}
	}

	sum := 0
	for _, v := range parts {
		result := execute_workflow(v, worflow)
		if result {
			sum += v.s
			sum += v.a
			sum += v.m
			sum += v.x
		}
	}

	assert.Equal(t, expected, sum)
}

func TestMainPart2(t *testing.T) {
	// your test code here

	expected := 167409079868000

	f, err := os.Open("inputTest.txt")

	if err != nil {
		t.Errorf("Error reading input file: %s", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	worflow := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, "{")

		if len(ls[0]) > 1 {
			nls := strings.Split(ls[1][:len(ls[1])-1], ",")
			worflow[ls[0]] = nls
		}
	}

	ranges := map[string][]int{}
	ranges["x"] = []int{1, 4000}
	ranges["m"] = []int{1, 4000}
	ranges["a"] = []int{1, 4000}
	ranges["s"] = []int{1, 4000}
	sum := count(ranges, "in", worflow)

	fmt.Printf("Sum: %d\n", sum)

	assert.Equal(t, expected, sum)
}
