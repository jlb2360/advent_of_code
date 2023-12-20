package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    nl "day1/num_list"
)

func find_word_num(line *string, word *string, list *nl.List){
    // First we want to find if the string exists
    occurance := strings.Contains(*line, *word)
    var st string
    
    if *word == "one"{
        st = "1"
    } else if *word == "two"{
        st = "2"
    }else if *word == "two"{
        st = "2"
    }else if *word == "three"{
        st = "3"
    }else if *word == "four"{
        st = "4"
    }else if *word == "five"{
        st = "5"
    }else if *word == "six"{
        st = "6"
    }else if *word == "seven"{
        st = "7"
    }else if *word == "eight"{
        st = "8"
    }else if *word == "nine"{
        st = "9"
    }



    if occurance == true {
        last_index := strings.LastIndex(*line, *word)
        first_index := strings.Index(*line, *word)
        if last_index == first_index{
            numb := nl.NumString{Num: st, Index: last_index}
            list.Insert(numb)
        } else {
            numb := nl.NumString{Num: st, Index: first_index}
            list.Insert(numb)
            numb = nl.NumString{Num: st, Index: last_index}
            list.Insert(numb)
        }
    }

}

func get_calib(line *string) int{
    num_in_word := nl.List{}

    var numbers [9]string = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    // first we want to check for any words that spell out numbers
    for _, word := range numbers {
        find_word_num(line, &word, &num_in_word)
    }


    for i, c := range *line{
        if _, err := strconv.Atoi(string(c)); err == nil{
            numb := nl.NumString{Num: string(c), Index: i}
            num_in_word.Insert(numb)
        }
    }


    n := num_in_word.GrabNum()
    fmt.Printf("The total number for %s is %v \n", *line, n)
    return n
}

func main(){
    f, err := os.Open("input-4.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()
    scanner := bufio.NewScanner(f)

    var sum int = 0

    for scanner.Scan() {
        line := scanner.Text()
        num := get_calib(&line)
        sum = sum + num
    }

    if err := scanner.Err(); err != nil{
        log.Fatal(err)
    }

    fmt.Printf("Calibration is: %v", sum)


}
