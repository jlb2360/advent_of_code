package main

import (
	"bufio"
	b "day15/box_map"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func HashAlgorithm(input string) int {
	current_value := 0
	for _, c := range input {
		//get hash value of c
		current_value += int(byte(c))
		current_value *= 17
		current_value %= 256
	}
	return current_value
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	var label string
	var command string
	var power int

	boxMap := b.BoxList{}
	boxMap.Construct()
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		fmt.Println(line)
		for _, test_instruction := range line {
			sum += HashAlgorithm(test_instruction)
			sl := strings.Split(test_instruction, "=")
			if len(sl) == 0 {
				fmt.Println(sl)
			}
			if len(sl) == 2 {
				label = sl[0]
				command = "="
				power, err = strconv.Atoi(string(sl[1]))
				if len(sl) > 2 {
					fmt.Println(sl)
				}

				if err != nil {
					log.Fatal(err)
				}
			} else {
				sl := strings.Split(test_instruction, "-")
				if len(sl) > 2 {
					fmt.Println(sl)
				}
				if len(sl) == 2 {
					label = sl[0]
					command = "-"
				}
			}

			l := b.Lense{label, power}

			hNum := get_hash_value(label)

			if command == "=" {
				boxMap.PerformEquals(hNum, l)
			} else if command == "-" {
				boxMap.PerformMinus(hNum, l)
			}
		}
	}

	log.Printf("sum: %d", sum)
	sum2 := boxMap.Sum()
	log.Printf("sum2: %d", sum2)

}

func get_hash_value(data string) int {
	result := 0
	for d := range data {
		result = ((result + int(data[d])) * 17) % 256
	}
	return result
}
