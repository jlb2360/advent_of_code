package main

import (
	"bufio"
	q "day22/queue"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestMainPart1(t *testing.T) {
	f, err := os.Open("inputTest.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	bricks := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "~", ",", -1)
		ls := strings.Split(line, ",")
		l := []int{}
		for _, s := range ls {
			num, _ := strconv.Atoi(s)
			l = append(l, num)
		}
		bricks = append(bricks, l)
	}

	bricks = sortBricks(bricks)

	for _, brick := range bricks {
		fmt.Printf("%v\n", brick)
	}

	for i := 0; i < len(bricks); i++ {
		brick := bricks[i]
		max_z := 1
		for j := 0; j < i; j++ {
			b_check := bricks[j]
			if (max(b_check[0], brick[0]) <= min(b_check[3], brick[3])) && (max(b_check[1], brick[1]) <= min(b_check[4], brick[4])) {
				max_z = max(max_z, b_check[5]+1)
			}
		}

		brick[5] -= brick[2] - max_z
		brick[2] = max_z
	}

	println("")

	for _, brick := range bricks {
		fmt.Printf("%v\n", brick)
	}

	k_sups_i := map[int][]int{}
	for i, _ := range bricks {
		k_sups_i[i] = []int{}
	}
	i_sups_k := map[int][]int{}
	for i, _ := range bricks {
		i_sups_k[i] = []int{}
	}

	for j := 0; j < len(bricks); j++ {
		upper := bricks[j]
		for i := 0; i < j; i++ {
			lower := bricks[i]
			if (max(lower[0], upper[0]) <= min(lower[3], upper[3])) && (max(lower[1], upper[1]) <= min(lower[4], upper[4])) && (upper[2] == lower[5]+1) {
				k_sups_i[i] = append(k_sups_i[i], j)
				i_sups_k[j] = append(i_sups_k[j], i)
			}
		}
	}

	total := 0

	for i := 0; i < len(bricks); i++ {
		supported := true
		for _, j := range k_sups_i[i] {
			if len(i_sups_k[j]) < 2 {
				supported = false
			}
		}
		if supported {
			total += 1
		}
	}

	fmt.Printf("Total: %v\n", total)
}

func TestMainPart2(t *testing.T) {
	f, err := os.Open("inputTest.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	bricks := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, "~", ",", -1)
		ls := strings.Split(line, ",")
		l := []int{}
		for _, s := range ls {
			num, _ := strconv.Atoi(s)
			l = append(l, num)
		}
		bricks = append(bricks, l)
	}

	bricks = sortBricks(bricks)

	for _, brick := range bricks {
		fmt.Printf("%v\n", brick)
	}

	for i := 0; i < len(bricks); i++ {
		brick := bricks[i]
		max_z := 1
		for j := 0; j < i; j++ {
			b_check := bricks[j]
			if (max(b_check[0], brick[0]) <= min(b_check[3], brick[3])) && (max(b_check[1], brick[1]) <= min(b_check[4], brick[4])) {
				max_z = max(max_z, b_check[5]+1)
			}
		}

		brick[5] -= brick[2] - max_z
		brick[2] = max_z
	}

	println("")

	for _, brick := range bricks {
		fmt.Printf("%v\n", brick)
	}

	k_sups_i := map[int][]int{}
	for i, _ := range bricks {
		k_sups_i[i] = []int{}
	}
	i_sups_k := map[int][]int{}
	for i, _ := range bricks {
		i_sups_k[i] = []int{}
	}

	for j := 0; j < len(bricks); j++ {
		upper := bricks[j]
		for i := 0; i < j; i++ {
			lower := bricks[i]
			if (max(lower[0], upper[0]) <= min(lower[3], upper[3])) && (max(lower[1], upper[1]) <= min(lower[4], upper[4])) && (upper[2] == lower[5]+1) {
				k_sups_i[i] = append(k_sups_i[i], j)
				i_sups_k[j] = append(i_sups_k[j], i)
			}
		}
	}

	total := 0

	queue := q.Queue{}

	for i := 0; i < len(bricks); i++ {
		falling := map[int]bool{}
		falling[i] = true
		for _, j := range k_sups_i[i] {
			if len(i_sups_k[j]) == 1 {
				queue.Enqueue(j)
				falling[j] = true
			}
		}

		for queue.Length > 0 {
			j := queue.Dequeue()
			for _, k := range k_sups_i[j] {
				if !falling[k] {
					found_subset := true
					for _, l := range i_sups_k[k] {
						if !falling[l] {
							found_subset = false
						}
					}

					if found_subset {
						queue.Enqueue(k)
						falling[k] = true
					}
				}
			}
		}

		total += len(falling) - 1
	}

	fmt.Printf("Total: %v\n", total)
}
