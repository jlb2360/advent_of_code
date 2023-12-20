package main

import (
	"bufio"
	m "day3/mat"
	"log"
	"os"
	"strconv"
)


func main(){
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    var id bool
    var is bool

    defer f.Close()
    scanner := bufio.NewScanner(f)

    mat := m.Mat{}

    j := 1
    for scanner.Scan() {
        line := scanner.Text()
        mat.InsertRow()
        for _, c := range line {
            if string(c) == "."{
                is = false   
            } else {
                is = true
            }


            if _, err := strconv.Atoi(string(c)); err == nil{
                id = true
                is = false
            } else {
                id = false
            }

            if string(c) == "."{
                is = false   
            } 
            elem := m.Element{Entry: string(c), IsDigit: id, IsShape: is}
            mat.InsertElement(elem, j)
        }
        j = j + 1
    }

    mat.GetSumParts()
    mat.SumGears()
}
