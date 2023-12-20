package main

import (
	"bufio"
	l "day10/list"
	la "day10/list/answer"
	"fmt"
	"log"
	"os"
)


func main(){
    f, err := os.Open("input.txt")
    if err != nil{
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    mat := l.LoL{}

    for scanner.Scan(){
        line := scanner.Text()
        list := l.List{}
        for _, c := range line{
            list.Insert_back(string(c))
        }
        mat.Insert(list)
    }


    s_r, s_c := mat.FindS() 
    for i:=-1; i<2; i++{
        for j:=-1; j<2; j++{
            pipe_around := mat.GetIdx(s_r+i, s_c+j)
            fmt.Printf("%s", pipe_around)
        }
        fmt.Printf("\n")
    }

    front_lis := la.List{}
    
    d := 0
    pos := la.Answer{Distance: 0, Pipe: "-", Col: s_c, Row: s_r}
    dir := "right"
    for pos.Pipe != "S"{
        if d != 0{
        front_lis.Insert_back(pos)
    }
        s_r, s_c, dir = mat.NextPipe(pos.Pipe, pos.Row, pos.Col, dir)
        if s_r == -1{
            fmt.Printf("%v, %v, %s", pos.Row, pos.Col, pos.Pipe)
            os.Exit(1)
        }
        d += 1
        pip := mat.GetIdx(s_r, s_c)
        pos = la.Answer{Distance: d, Pipe: pip, Col: s_c, Row: s_r}
    }

    front_lis.PrintList()


    fmt.Println("-------------------------------------------------")
    fmt.Println("Now Back List")
    fmt.Println("-------------------------------------------------")

    s_r, s_c = mat.FindS() 
    back_list := la.List{}
    
    d = 0
    pos = la.Answer{Distance: 0, Pipe: "-", Col: s_c, Row: s_r}
    dir = "left"
    for pos.Pipe != "S"{
        if d != 0{
        back_list.Insert_front(pos)
        }
        s_r, s_c, dir = mat.NextPipe(pos.Pipe, pos.Row, pos.Col, dir)
        if s_r == -1{
            fmt.Printf("%v, %v, %s", pos.Row, pos.Col, pos.Pipe)
            os.Exit(1)
        }
        d += 1
        pip := mat.GetIdx(s_r, s_c)
        pos = la.Answer{Distance: d, Pipe: pip, Col: s_c, Row: s_r}
    }

    back_list.PrintList()


    answer_list := front_lis.MinList(&front_lis, &back_list)

    
    fmt.Println("-------------------------------------------------")
    fmt.Println("Now Answer List")
    fmt.Println("-------------------------------------------------")

    answer_list.PrintList()


    m := answer_list.GetMax()

    fmt.Printf("Answer is %v\n", m.Distance)

}
