package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := make([][]string, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		input = append(input, line)
	}

	platform := Platform{Platform: input}

	platform.MoveAll()

	fmt.Printf("load: %v\n", platform.CalcLoad())

	platform = Platform{Platform: input}

	load := platform.CycleN(1000000000)

	fmt.Printf("load: %v\n", load)
}
