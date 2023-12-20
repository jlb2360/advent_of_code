package main

import (
	"bufio"
	l "day6/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calc_wins(d int, t int) int{
    fmt.Printf("Distance to Cover: %v\n Time to use: %v\n", d, t)
    wins := 0
    for i:=0; i<t; i++{
        //i is ms of charge
        speed := i // mm/ms
        dt := t-i //ms
        dist := speed*dt
        if dist>d{
            wins = wins + 1
        }

    }
    return wins
}


func main(){
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)

    times := l.List{}
    powers := 1
    var dis string
    var ti  string

    i := 0
    for scanner.Scan(){
        line := scanner.Text()
        if i == 0{
            ts := strings.Split(line, " ")
            for _, t := range ts{
                t, err := strconv.Atoi(t)
                if err == nil{
                    times.Insert(t)
                }
            }
        }

        if i == 0{
            ts := strings.Split(line, " ")
            for _, time_l := range ts{
                _, err := strconv.Atoi(time_l)
                if err == nil{
                    ti = ti + time_l
                }
            }
        }

        if i == 1{
            ds := strings.Split(line, " ")
            j := 0
            for _, d := range ds{
                d, err := strconv.Atoi(d)
                if err == nil{
                    t := times.GetIdx(j)
                    wins := calc_wins(d,t)
                    powers = powers*wins
                    j = j + 1
                }
            }
        }

        if i == 1 {
            ds := strings.Split(line, " ")
            for _, d_l := range ds{
                _, err := strconv.Atoi(d_l)
                if err == nil {
                    dis = dis + d_l
                }
            }
        }
        i = i + 1
    }

    d, err := strconv.Atoi(dis)
    if err != nil{
        log.Fatal(err)
    }
    t, err := strconv.Atoi(ti)
    if err != nil{
        log.Fatal(err)
    }
    wins := calc_wins(d, t)

    fmt.Printf("Actual wins %v\n",wins)

    fmt.Printf("Powers from wins %v\n", powers)
}
