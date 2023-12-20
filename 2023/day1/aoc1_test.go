package main

import "testing"

func TestWordDigit1(t *testing.T){
    test_word := "two1nine"
    num := get_calib(&test_word) 

    if num != 29 {
        t.Fatalf("String %s did not give 29 instead we got %v", test_word, num)
    }
}
