package main

import (
	"bufio"
	l "day9/list"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
    f, err := os.Open("input.txt")
    if err != nil{
        log.Fatal(err)
    }
    sum_for := 0
    sum_back := 0

    scanner := bufio.NewScanner(f)

    for scanner.Scan(){
        line := scanner.Text()
        list_of_list := l.LoL{}
        num_list := l.List{}

        sl := strings.Split(line, " ")

        for _, c := range sl{
            num, err := strconv.Atoi(string(c))
            if err == nil{
                num_list.Insert(num)
            }
        }

        list_of_list.Insert(num_list)
        check := false
        new_lis := num_list
        for check == false{
            new_lis = new_lis.CalcDif()
            list_of_list.Insert(new_lis)
            
            check = new_lis.CheckZeros()
        }

        n_for := list_of_list.Cald_dif()
        n_back := list_of_list.Cald_back_dif()

        sum_for = sum_for+n_for
        sum_back = sum_back+n_back
    }

    fmt.Printf("Sum is %v\n",sum_for)
    fmt.Printf("Sum is %v\n",sum_back)
}
