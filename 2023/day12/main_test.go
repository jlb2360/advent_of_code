package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssesment(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		records  []int
	}{
		{
			expected: 1,
			input:    "???.###",
			records:  []int{1, 1, 3},
		},
		{
			expected: 4,
			input:    ".??..??...?##.",
			records:  []int{1, 1, 3},
		},
		{
			expected: 1,
			input:    "?#?#?#?#?#?#?#?",
			records:  []int{1, 3, 1, 6},
		},
		{
			expected: 1,
			input:    "????.#...#...",
			records:  []int{4, 1, 1},
		},
		{
			expected: 4,
			input:    "????.######..#####.",
			records:  []int{1, 6, 5},
		},
		{
			expected: 10,
			input:    "?###????????",
			records:  []int{3, 2, 1},
		},
	}

	for _, test := range tests {
		cache := make(map[string]int)

		assert.Equal(t, test.expected, Count(test.input, test.records, cache))
	}
}

func TestAssesment2(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		records  []int
	}{
		{
			expected: 1,
			input:    "???.###",
			records:  []int{1, 1, 3},
		},
		{
			expected: 16384,
			input:    ".??..??...?##.",
			records:  []int{1, 1, 3},
		},
		{
			expected: 1,
			input:    "?#?#?#?#?#?#?#?",
			records:  []int{1, 3, 1, 6},
		},
		{
			expected: 16,
			input:    "????.#...#...",
			records:  []int{4, 1, 1},
		},
		{
			expected: 2500,
			input:    "????.######..#####.",
			records:  []int{1, 6, 5},
		},
		{
			expected: 506250,
			input:    "?###????????",
			records:  []int{3, 2, 1},
		},
	}

	for _, test := range tests {

		cache := make(map[string]int)

		parts := strings.Repeat(test.input+"?", 5)
		parts = parts[:len(parts)-1]
		fmt.Printf("parts: %s\n", parts)

		rec := []int{}
		for i := 0; i < 5; i++ {
			rec = append(rec, test.records...)
		}

		fmt.Printf("rec: %v\n", rec)

		assert.Equal(t, test.expected, Count(parts, rec, cache))
	}
}
