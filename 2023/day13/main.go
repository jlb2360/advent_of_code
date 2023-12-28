package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func transpose(input [][]string) [][]string {
	output := make([][]string, len(input[0]))
	for i := 0; i < len(input[0]); i++ {
		output[i] = make([]string, len(input))
		for j := 0; j < len(input); j++ {
			output[i][j] = input[j][i]
		}
	}
	return output
}

func reverse(input [][]string) [][]string {
	output := make([][]string, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = make([]string, len(input[len(input)-1-i]))
		for j := 0; j < len(input[len(input)-1-i]); j++ {
			output[i][j] = input[len(input)-1-i][j]
		}
	}
	return output
}

func FindReflection(input [][]string) ([]int, []int) {
	// First we want to check if there is a vertical reflection
	// transpose the input
	inputT := transpose(input)
	result_h := []int{}
	result_v := []int{}
	for i := 0; i < len(inputT)-1; i++ {
		check_len := min(len(inputT[0:i+1]), len(inputT[i+1:]))
		arr1 := inputT[i+1-check_len : i+1]
		arr2 := inputT[i+1 : i+1+check_len]
		arr2 = reverse(arr2)
		if reflect.DeepEqual(arr1, arr2) {
			result_v = append(result_v, (i + 1))
		}
	}

	for i := 0; i < len(input)-1; i++ {
		check_len := min(len(input[0:i+1]), len(input[i+1:]))
		arr1 := input[i+1-check_len : i+1]
		arr2 := input[i+1 : i+1+check_len]
		arr2 = reverse(arr2)
		if reflect.DeepEqual(arr1, arr2) {
			result_h = append(result_h, (i + 1))
		}
	}

	return result_h, result_v

}

func Check(input []int, first []int) bool {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(first); j++ {
			if input[i] != first[j] {
				return false
			}
		}
	}
	return false
}

func NewReflection(first_num_h []int, first_num_v []int, input [][]string) (int, int) {
	new_num_h := 0
	new_num_v := 0
	old_num_h := Sum_h(first_num_h)
	old_num_v := Sum_v(first_num_v)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "." {
				input[i][j] = "#"
			} else {
				input[i][j] = "."
			}

			ref_num_h, ref_num_v := FindReflection(input)
			sum_num_h := Sum_h(ref_num_h)
			sum_num_v := Sum_v(ref_num_v)
			fmt.Printf("first_num_h: %v, first_num_v: %v\n", first_num_h, first_num_v)
			fmt.Printf("ref_num_h: %v, ref_num_v: %v\n", ref_num_h, ref_num_v)
			if (sum_num_h == old_num_h && sum_num_v == old_num_v) || ((sum_num_h == 0) && (sum_num_v == 0)) {
				if input[i][j] == "." {
					input[i][j] = "#"
				} else {
					input[i][j] = "."
				}
			} else {

				for k := 0; k < len(ref_num_h); k++ {
					for m := 0; m < len(first_num_h); m++ {
						if ref_num_h[k] == first_num_h[m] {
							sum_num_h -= first_num_h[m] * 100
						}
					}
				}

				for k := 0; k < len(ref_num_v); k++ {
					for m := 0; m < len(first_num_v); m++ {
						if ref_num_v[k] == first_num_v[m] {
							sum_num_v -= first_num_v[m]
						}
					}
				}
				return sum_num_h, sum_num_v

			}
		}
	}

	return new_num_h, new_num_v
}

func Sum_v(input []int) int {
	sum := 0
	for _, v := range input {
		sum += v
	}
	return sum
}

func Sum_h(input []int) int {
	sum := 0
	for _, v := range input {
		sum += v * 100
	}
	return sum
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := make([][]string, 0)
	check := make([]string, 0)
	sum := 0
	sum2 := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		if reflect.DeepEqual(line, check) {
			num_h, num_v := FindReflection(input)
			sum_h := Sum_h(num_h)
			sum_v := Sum_v(num_v)
			sum += sum_h + sum_v

			new_num_h, new_num_v := NewReflection(num_h, num_v, input)
			sum2 += new_num_h + new_num_v

			input = make([][]string, 0)
		} else {
			input = append(input, line)
		}
	}

	num_h, num_v := FindReflection(input)
	sum_h := Sum_h(num_h)
	sum_v := Sum_v(num_v)
	sum += sum_h + sum_v

	new_num_h, new_num_v := NewReflection(num_h, num_v, input)
	sum2 += new_num_h + new_num_v
	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("sum2: %v\n", sum2)
}
