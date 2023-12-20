package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
    l "day2/list"
)

func check_balls(line *string) (bool, int){
    ss := strings.SplitAfter(*line, ":")
    var id string
    for _, b := range ss{
        if strings.Contains(b, ":"){
            for _, c := range b {
                if _, err := strconv.Atoi(string(c)); err == nil{
                   id = id + string(c)
                }
            }
        } else {
            list_of_balls := strings.SplitAfter(b, ";")
            for _, lb := range list_of_balls{
                
                balls := strings.SplitAfter(lb, ",")
                for _, ball := range balls{
                    var num string
                    for _, c := range ball {
                        if _, err := strconv.Atoi(string(c)); err == nil{
                            num = num + string(c)
                        }
                    }

                    n, err := strconv.Atoi(num)
                    if err != nil {
                        log.Fatal(err)
                    }

                    color := strings.SplitAfter(ball, " ")[2]
                    if strings.Contains(color, ","){
                        color = color[:len(color)-1]
                    }
                    if strings.Contains(color, ";"){
                        color = color[:len(color)-1]
                    }


                    switch color {
                    case "red":{
                            if n > 12{
                                return false, 0
                            }
                        }
                    case "green": {
                            if n > 13{
                                return false, 0
                            }
                        }
                    case "blue": {
                            if n > 14{
                                return false, 0
                            }
                        }
                    }
                }

            }
        }
    }
    id_num, err := strconv.Atoi(id)
    if err != nil {
        log.Fatal(err)
    }
   return true, id_num
}

func get_powers(line *string) int{
    ss := strings.SplitAfter(*line, ":")
    reds := l.List{}
    greens := l.List{}
    blues := l.List{}
    var id string
    for _, b := range ss{
        if strings.Contains(b, ":"){
            for _, c := range b {
                if _, err := strconv.Atoi(string(c)); err == nil{
                   id = id + string(c)
                }
            }
        } else {
            list_of_balls := strings.SplitAfter(b, ";")
            for _, lb := range list_of_balls{
                
                balls := strings.SplitAfter(lb, ",")
                for _, ball := range balls{
                    var num string
                    for _, c := range ball {
                        if _, err := strconv.Atoi(string(c)); err == nil{
                            num = num + string(c)
                        }
                    }

                    n, err := strconv.Atoi(num)
                    if err != nil {
                        log.Fatal(err)
                    }

                    color := strings.SplitAfter(ball, " ")[2]
                    if strings.Contains(color, ","){
                        color = color[:len(color)-1]
                    }
                    if strings.Contains(color, ";"){
                        color = color[:len(color)-1]
                    }


                    switch color {
                    case "red":
                        reds.Insert(n)
                    case "green": 
                        greens.Insert(n)
                    case "blue": 
                        blues.Insert(n)
                        }
                    }
                }

            }
        }
    

    powers := reds.GetMax()*greens.GetMax()*blues.GetMax()
    return powers
}
func main(){
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()
    scanner := bufio.NewScanner(f)

    var sum int = 0
    var powers int = 0

    for scanner.Scan() {
        line := scanner.Text()
        occ, num := check_balls(&line)
        if occ == true {
            sum = sum + num
        }
        power := get_powers(&line)
        powers = powers+power
    }

    if err := scanner.Err(); err != nil{
        log.Fatal(err)
    }

    fmt.Printf("sum of ids is: %v", sum)
    fmt.Printf("sum of powers is: %v", powers)



}
