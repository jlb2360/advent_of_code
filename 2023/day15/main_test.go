package main

import (
	b "day15/box_map"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashAlgorithm(t *testing.T) {

	test_string := "HASH"
	expected := 52

	result := HashAlgorithm(test_string)

	assert.Equal(t, expected, result, "HashAlgorithm(%s) should be %d", test_string, expected)
}

func TestPart2(t *testing.T) {

	test_instructions := []string{"rn=1", "cm-", "qp=3", "cm=2", "qp-", "pc=4", "ot=9", "ab=5", "pc-", "pc=6", "ot=7"}
	expected := 145

	var label string
	var command string
	var power int
	var err error

	boxMap := b.BoxList{}
	boxMap.Construct()

	for _, test_instruction := range test_instructions {
		sl := strings.Split(test_instruction, "=")
		if len(sl) == 2 {
			label = sl[0]
			command = "="
			power, err = strconv.Atoi(string(sl[1]))

			if err != nil {
				log.Fatal(err)
			}
		} else {
			sl := strings.Split(test_instruction, "-")
			if len(sl) == 2 {
				label = sl[0]
				command = "-"
			}
		}

		l := b.Lense{label, power}

		hNum := HashAlgorithm(label)

		if command == "=" {
			boxMap.PerformEquals(hNum, l)
		} else if command == "-" {
			boxMap.PerformMinus(hNum, l)
		}

		boxMap.PrintBox()
		fmt.Printf("\n ========= \n")
	}

	result := boxMap.Sum()

	assert.Equal(t, expected, result, "Part2 should be %d", expected)
}
