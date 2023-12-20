package main

import (
    "os"
    "bufio"
	l "day4/list"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func check_winner(line *string) (int, int){
    heads := strings.Split(*line, ":")
    lists := strings.Split(heads[1], "|")

    winners := l.List{}
    ours := l.List{}
    num := ""

    for _, c := range lists[0]{
        if string(c) != " "{
            if _, err := strconv.Atoi(string(c)); err == nil{
                num = num + string(c)
            }
        } else {
            if num != ""{
                n, err := strconv.Atoi(num)
                if err == nil {
                    winners.Insert(n)
                }
            }
            num = ""
        }
    }

    num = ""

    for _, c := range lists[1]{
        if string(c) != " "{
            if _, err := strconv.Atoi(string(c)); err == nil{
                num = num + string(c)
            }
        } else {
            if num != ""{
                n, err := strconv.Atoi(num)
                if err == nil {
                    ours.Insert(n)
                }
            }
            num = ""
        }
    }

    
    if num != ""{
        n, err := strconv.Atoi(num)
        if err == nil {
            ours.Insert(n)
        }
    }

    points := 0
    match := 0

    for i := 1; i<=ours.Length; i++{
        our_num, err := ours.GrabIdx(i)
        if err != nil {
            log.Fatal(err)
        }

        check := winners.CheckIn(our_num)
        if check == true{
            match = match + 1
            if points == 0{
                points = points + 1
            } else if points == 1 {
                points = points + 1
            } else{ points = points*2 }
        
        }

    }




    return points, match
}


func main(){

    f, err := os.Open("input.txt")
    if err != nil {
       log.Fatal(err) 
    }

    defer f.Close()
    scanner := bufio.NewScanner(f)
    sum := 0
    tot_cards := 0
    var future_copies [26]int
    for i:= 0; i<len(future_copies); i++{
        future_copies[i] = 1
    }

    for scanner.Scan(){
        line := scanner.Text()
        num, match := check_winner(&line)
        sum = sum + num
        tot_cards = tot_cards + future_copies[0]
        copies := future_copies[0]
        for i:= 0; i<len(future_copies)-1; i++{
            future_copies[i] = future_copies[i+1]
        }

        future_copies[len(future_copies)-1] = 1
        



        for j:=0; j<copies; j++{
            for i:=0; i<match; i++{
                future_copies[i] = future_copies[i] + 1 
            }
        }
    }

    fmt.Printf("The sum is: %v", sum)
    fmt.Printf("The total cards is: %v", tot_cards)


}
