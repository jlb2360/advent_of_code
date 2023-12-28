package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Count(parts string, records []int, cache map[string]int) int {

	if parts == "" {
		if len(records) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(records) == 0 {
		for _, c := range parts {
			if c == '#' {
				return 0
			}
		}

		return 1
	}

	key := parts
	for _, r := range records {
		key += strconv.Itoa(r)
	}

	val, ok := cache[key]
	if ok {
		return val
	}

	result := 0
	if (parts[0] == '.') || (parts[0] == '?') {
		result += Count(parts[1:], records, cache)
	}

	if (parts[0] == '?') || (parts[0] == '#') {
		if records[0] <= len(parts) {
			for i := 0; i < records[0]; i++ {
				if parts[i] == '.' {
					return result
				}
			}

			if records[0] == len(parts) {
				return result + Count(parts[records[0]:], records[1:], cache)
			}

			if parts[records[0]] != '#' {
				result += Count(parts[records[0]+1:], records[1:], cache)
			}
		}
	}

	cache[key] = result

	return result
}

func s2i(ss []string) []int {
	var ints []int
	for _, s := range ss {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}
	}

	return ints
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	sum := 0
	sum2 := 0

	i := 0
	cache := make(map[string]int)
	for scanner.Scan() {
		prog := float64(i) / 1000.0
		fmt.Printf("%v\n", prog)
		text := scanner.Text()
		if text == "" {
			continue
		}

		sl := strings.Split(text, " ")
		parts := sl[0]
		records := s2i(strings.Split(sl[1], ","))

		cache = make(map[string]int)

		sum += Count(parts, records, cache)

		//fmt.Printf("parts: %s\n", parts)

		parts = strings.Repeat(parts+"?", 5)
		parts = parts[:len(parts)-1]

		rec := []int{}
		for i := 0; i < 5; i++ {
			rec = append(rec, records...)
		}

		sum2 += Count(parts, rec, cache)

		i++
	}

	fmt.Println(sum)
	fmt.Println(sum2)

}
