package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatternPart1(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test1.txt",
			expected:  5,
		},
		{
			inputFile: "input_test2.txt",
			expected:  400,
		},
	}

	for _, in := range tests {
		f, err := os.Open(in.inputFile)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		input := make([][]string, 0)
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), "")
			input = append(input, line)
		}

		fmt.Printf("input: %v\n", input)

		ref_h, ref_v := FindReflection(input)
		num_h := Sum_h(ref_h)
		num_v := Sum_v(ref_v)
		num := num_h + num_v

		assert.Equal(t, in.expected, num)
	}
}

func TestPatternPart2(t *testing.T) {
	tests := []struct {
		inputFile string
		expected  int
	}{
		{
			inputFile: "input_test1.txt",
			expected:  300,
		},
		{
			inputFile: "input_test2.txt",
			expected:  100,
		},
	}

	for _, in := range tests {
		f, err := os.Open(in.inputFile)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		input := make([][]string, 0)
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), "")
			input = append(input, line)
		}

		first_num_h, first_num_v := FindReflection(input)
		fmt.Printf("first_num_h: %v, first_num_v: %v\n", first_num_h, first_num_v)

		sum_num_h, sum_num_v := NewReflection(first_num_h, first_num_v, input)

		new_num := sum_num_h + sum_num_v

		assert.Equal(t, in.expected, new_num)
	}
}
