package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain1(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test1.txt",
			expected:  136,
		},
	}

	for _, test := range tests {
		f, err := os.Open(test.inputFile)
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

		platform.PrintPlat()

		platform.MoveAll()

		fmt.Printf("\nafter move\n")
		platform.PrintPlat()

		assert.Equal(t, test.expected, platform.CalcLoad())
	}
}

func TestMain2(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test1.txt",
			expected:  64,
		},
	}

	for _, test := range tests {
		f, err := os.Open(test.inputFile)
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

		load := platform.CycleN(1000000000)

		assert.Equal(t, test.expected, load)
	}
}
