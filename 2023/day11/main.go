package main


import (
	"bufio"
	l "day11/utilities/mat"
	"log"
	"os"
    "fmt"
)




func main(){
    f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(f)
    data := l.Mat{}

    for scanner.Scan(){
        line := scanner.Text()
        lis := l.List{}
        var wy int64 = 1000000
        for _, c := range line{
            if string(c) == "#"{
                wy = 1
            }
        }
        for _, c := range line{
            lis.Insert(string(c), wy, 1)
        }
        data.Insert(&lis)
    }

    data.CheckStarCol(1000000)


    gal_rows, gal_cols := data.GetGals()
    sum := 0

    for i, r:= range gal_rows{
        prog := float64(i)/float64(len(gal_rows))
        fmt.Printf("Progress: %v\n", prog)
        data.Reset()
        init := l.Corridinent{Row: r, Col: gal_cols[i]}
        data.SetInit(r, gal_cols[i])
        data.AStar(init)
        for j:=i+1; j<len(gal_cols); j++{
            num := data.GetDis(gal_rows[j], gal_cols[j])
            sum += int(num)
        }
    }


    fmt.Printf("Sum: %v\n", sum)
    

}
