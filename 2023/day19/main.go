package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

func execute_workflow(part Part, workflows map[string][]string) bool {
	workflow_name := "in"

	working := true
	for working {
		fmt.Printf("%s -> ", workflow_name)
		if workflow_name == "R" {
			fmt.Println("Rejected")
			working = false
			return false
		}
		if workflow_name == "A" {
			fmt.Println("Accepted")
			working = false
			return true
		}
		workflow := workflows[workflow_name]
		for i := 0; i < len(workflow); i++ {
			check := workflow[i]
			if check == "A" {
				fmt.Println("Accepted")
				working = false
				return true
			} else if check == "R" {
				fmt.Println("Rejected")
				working = false
				return false
			}

			ls := strings.Split(check, ":")
			num, _ := strconv.Atoi(ls[0][2:])
			quality := string(ls[0][0])
			condition := string(ls[0][1])

			if len(ls) == 1 {
				workflow_name = ls[0]
				continue
			}
			switch quality {
			case "x":
				if condition == ">" {
					if part.x > num {
						workflow_name = ls[1]
						i = len(workflow)
						continue
					}
				} else if condition == "<" {
					if part.x < num {
						workflow_name = ls[1]
						i = len(workflow)
						continue
					}
				}
			case "m":
				if condition == ">" {
					if part.m > num {
						workflow_name = ls[1]
						i = len(workflow)
						continue
					}
				} else if condition == "<" {
					if part.m < num {
						workflow_name = ls[1]
						i = len(workflow)
						continue
					}
				}
			case "a":
				if condition == ">" {
					if part.a > num {
						workflow_name = ls[1]
						i = len(workflow)
						continue
					}
				} else if condition == "<" {
					if part.a < num {
						workflow_name = ls[1]
						i = len(workflow)
						continue
					}
				}
			case "s":
				if condition == ">" {
					if part.s > num {
						workflow_name = ls[1]
						i = len(workflow)
						continue
					}
				} else if condition == "<" {
					if part.s < num {
						workflow_name = ls[1]
						i = len(workflow)
						continue
					}
				}
			}
		}
	}

	return false
}

func count(ranges map[string][]int, name string, worksflows map[string][]string) int {
	var T, F []int

	if name == "R" {
		return 0
	} else if name == "A" {
		product := 1
		for _, v := range ranges {
			product *= v[1] - v[0] + 1
		}
		return product
	}

	total := 0

	workflow := worksflows[name]
	for i := 0; i < len(workflow); i++ {
		check := workflow[i]

		ls := strings.Split(check, ":")

		if len(ls) == 1 {
			total += count(ranges, ls[0], worksflows)
			continue
		}

		num, _ := strconv.Atoi(ls[0][2:])
		key := string(ls[0][0])
		condition := string(ls[0][1])

		lo, hi := ranges[key][0], ranges[key][1]
		if condition == "<" {
			T = []int{lo, num - 1}
			F = []int{num, hi}
		} else {
			T = []int{num + 1, hi}
			F = []int{lo, num}
		}

		if T[0] <= T[1] {
			copy := map[string][]int{}
			for k, v := range ranges {
				copy[k] = v
			}
			copy[key] = T
			total += count(copy, ls[1], worksflows)
		}

		if F[0] <= F[1] {
			ranges[key] = F
		} else {
			break
		}
	}

	return total
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

	fmt.Printf("Sum: %d\n", sum)
}

func part2() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
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
}
