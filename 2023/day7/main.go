package main

import (
	"bufio"
	q "day7/queue"
	q2 "day7/queueP2"
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

	queue := q.PriorityQueue{}
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, " ")
		bid, _ := strconv.Atoi(ls[1])
		queue.Enqueue(ls[0], bid)
	}

	sum := 0
	i := 1
	for queue.Length > 0 {
		_, bid, _ := queue.Dequeue()
		sum += bid * i
		i++
	}

	fmt.Printf("Part1: %v\n", sum)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	queue := q2.PriorityQueue{}
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, " ")
		bid, _ := strconv.Atoi(ls[1])
		queue.Enqueue(ls[0], bid)
	}

	sum := 0
	i := 1
	for queue.Length > 0 {
		_, bid, _ := queue.Dequeue()
		sum += bid * i
		i++
	}

	fmt.Printf("Part2: %v\n", sum)
}
