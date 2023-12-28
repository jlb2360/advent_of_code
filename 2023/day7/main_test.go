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

	queue := q.PriorityQueue{}
	m := 0
	for scanner.Scan() {
		if m == 5 {
			break
		}
		line := scanner.Text()
		ls := strings.Split(line, " ")
		fmt.Printf("ls: %v\n", ls)
		bid, _ := strconv.Atoi(ls[1])
		queue.Enqueue(ls[0], bid)
		m++
	}

	sum := 0
	i := 1
	for queue.Length > 0 {
		hand, bid, val := queue.Dequeue()
		fmt.Printf("hand: %v, bid: %v val: %v\n", hand, bid, val)
		sum += bid * i
		i++
	}

	fmt.Printf("sum: %v\n", sum)
	assert.Equal(t, 6440, sum)
}

func TestMainPart2(t *testing.T) {
	f, err := os.Open("inputTest.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	queue := q2.PriorityQueue{}
	for scanner.Scan() {
		line := scanner.Text()
		ls := strings.Split(line, " ")
		fmt.Printf("ls: %v\n", ls)
		bid, _ := strconv.Atoi(ls[1])
		queue.Enqueue(ls[0], bid)
	}

	sum := 0
	i := 1
	for queue.Length > 0 {
		hand, bid, val := queue.Dequeue()
		fmt.Printf("hand: %v, bid: %v val: %v\n", hand, bid, val)
		sum += bid * i
		i++
	}

	fmt.Printf("sum: %v\n", sum)
}
